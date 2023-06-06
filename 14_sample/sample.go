package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "flag"
)

func main() {

    percent := flag.Float64("p", 10.0, "percentage (of 100) of input to output")
	flag.Parse()
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
		randomNumber := rand.Float64() * 100
        if randomNumber < *percent {
            fmt.Println(scanner.Text())
        }
    }
}
