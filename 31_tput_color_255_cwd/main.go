package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
)

func colorByte(s string) uint8 {
	sum := sha256.Sum256([]byte(s))
	return uint8(sum[0])
}

func main() {
	fi, err := os.Stdin.Stat()
	hasStdin := err == nil && (fi.Mode()&os.ModeCharDevice) == 0

	// STDIN MODE: colorize each line independently
	if hasStdin {
		scanner := bufio.NewScanner(os.Stdin)
		// allow long lines (find can produce long paths)
		buf := make([]byte, 0, 64*1024)
		scanner.Buffer(buf, 1024*1024)

		for scanner.Scan() {
			line := scanner.Text()
			c := colorByte(line)
			fmt.Printf("\x1b[38;5;%dm%s\x1b[0m\n", c, line)
		}
		return
	}

	// NO STDIN: hash normalized working directory and print number
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(0)
		return
	}

	if abs, err := filepath.Abs(wd); err == nil {
		wd = abs
	}
	if real, err := filepath.EvalSymlinks(wd); err == nil {
		wd = real
	}

	fmt.Println(colorByte(wd))
}