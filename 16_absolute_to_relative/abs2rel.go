package main

import (
    "bufio"
    "fmt"
    "os"
//    "flag"
    "path/filepath"
)

func main() {

//    percent := flag.Float64("p", 10.0, "percentage (of 100) of input to output")
	//flag.Parse()
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {

		line := scanner.Text()
		path := filepath.FromSlash(line)

		// Get the current working directory.
		currentWorkingDirectory, _ := filepath.Abs(".")

		// Convert the path to a relative path.
		relativePath, _ := filepath.Rel(currentWorkingDirectory, path)

		// Print the relative path.
		fmt.Println(relativePath)

    }
}
