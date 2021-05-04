package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Thanks to this page for the original code:
// https://ebixio.com/blog/2012/03/05/editing-gedcom-files-with-vim
func main() {

	in := bufio.NewReader(os.Stdin)
	for {
		s, err := in.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		trimmed := strings.TrimSpace(s)
		level, err := strconv.Atoi(strings.Split(trimmed, " ")[0])
		fmt.Print(strings.Repeat("\t", level))
		fmt.Print(trimmed)
		fmt.Print("\n")
	}
}
