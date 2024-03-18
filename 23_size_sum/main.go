package main

import (
"bufio"
	"fmt"
//	"io/ioutil"
  "github.com/dustin/go-humanize"

	"os"
)

func main() {
	var totalSize int64

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		filename := scanner.Text()
		fileInfo, err := os.Stat(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			continue
		}
		if !fileInfo.IsDir() { // Ignore directories
			totalSize += fileInfo.Size()
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading standard input: %s\n", err)
	}

//	fmt.Printf("Total size: %d bytes\n", totalSize)
	  humanSize := humanize.Bytes(uint64(totalSize))
  fmt.Printf("%s\n", humanSize)

}
