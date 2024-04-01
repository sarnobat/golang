package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

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
		line := s

		maxLength := 80
		// 		indentLevel := 2
		if len(line) > maxLength {
			foldedLines := splitLongLineToArray(line, maxLength)
			fmt.Println(foldedLines[0])
			for _, foldedLine := range foldedLines[1:] {
				fmt.Println("  " + foldedLine)
			}
		} else {
			fmt.Println(line)
		}
	}
}

func splitLongLineToArray(str string, limit int) []string {
	var lines []string
	var currentLine string
	for _, char := range str {
		if len(currentLine) == limit {
			lines = append(lines, currentLine)
			currentLine = ""
		}
		currentLine += string(char)
	}
	// Add the last line segment (if not empty)
	if len(currentLine) > 0 {
		lines = append(lines, currentLine)
	}
	return lines
}
