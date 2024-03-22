package main

import (
    "fmt"
    "os"
    "github.com/pborman/getopt"
)

func main() {
    optName := getopt.StringLong("name", 'n', "", "Your name")
    optHelp := getopt.BoolLong("help", 0, "Help")
    getopt.Parse()

    if *optHelp {
        getopt.Usage()
        os.Exit(0)
    }

    fmt.Println("Hello " + *optName)
}
