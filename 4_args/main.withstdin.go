package main

import (
	"bufio"
	"fmt"
	"github.com/pborman/getopt"
	"io"
	"log"
	"os"
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

	in := bufio.NewReader(os.Stdin)
	for {
		s, err := in.ReadString('\n')
		if err != nil {
			// io.EOF is expected, anything else
			// should be handled/reported
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		// Do something with the line of text
		// in string variable s.
		_ = s
		fmt.Print("added: " + s)
	}

}
