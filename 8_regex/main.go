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

		exp := "^.*DOCUMENT_FREQUENCY_TOTAL.*\\s*([a-zA-Z0-9].*)\\s*$"
		r := regexp.MustCompile(exp)
		elem := r.FindStringSubmatch(s)

		if len(elem) == 0 {
			// no match
			continue
		}

		pathSegments[len(elem[1])] = elem[2]

		for i := 0; i < len(elem[1]); i++ {
			fmt.Print(pathSegments[i])
			fmt.Print("/")
		}
		fmt.Print(elem[2])

		fmt.Println()

	}

}
