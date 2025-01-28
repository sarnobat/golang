package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bbalet/stopwords"
)

func main() {
	// Create a scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)

	// Prompt user for input
	fmt.Println("Enter lines of text (type 'exit' to stop):")

	// Read each line from stdin
	for scanner.Scan() {
		line := scanner.Text()

		// Optionally, handle case when user types 'exit'
		if line == "exit" {
			break
		}

		// Convert the line to lowercase and remove stopwords
		// You can also specify a language (default is 'en' for English)
		cleanedLine := stopwords.CleanString(line, stopwords.SingleWord|stopwords.CaseSensitive)

		// If the cleaned line is empty after removing stopwords, don't print
		if cleanedLine != "" {
			fmt.Println(cleanedLine)
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

