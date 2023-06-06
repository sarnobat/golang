package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "flag"
)

func main() {
    // Create a new argument parser
//     parser := flag.NewParser()

    // Define an argument for the percentage
    percent := flag.Float64("p", 10.0, "percentage (of 100) of input to output")

    // Parse the arguments
//     args, err := parser.Parse()
//     if err != nil {
//         fmt.Println(err)
//         return
//     }

    // Read the lines from the standard input
    scanner := bufio.NewScanner(os.Stdin)


    // Print the lines that are less than the percentage
    for scanner.Scan() {
		// Generate a random number between 0 and 100
		randomNumber := rand.Float64() * 100

        if randomNumber < *percent {
            fmt.Println(scanner.Text())
        }
    }
}
