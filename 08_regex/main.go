package main

import (
	"bufio"
	"fmt"
	"regexp"
	"github.com/pborman/getopt"
	"io"
	"log"
	"os"
)

func main() {
	optName := getopt.StringLong("name", 'n', "Sridhar", "Your name")
	optHelp := getopt.BoolLong("help", 0, "Help")
	getopt.Parse()

	if *optHelp {
		getopt.Usage()
		os.Exit(0)
	}
	fmt.Println("name = " + *optName)

	in := bufio.NewReader(os.Stdin)
	for {
		line, err := in.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		_ = line
		fmt.Print("[debug] line: " + line)

		regex := "^\\s*([0-9]+)*\\s*DOCUMENT_FREQUENCY_TOTAL..(.*)\n"
		r := regexp.MustCompile(regex)
		elem := r.FindStringSubmatch(line)

		if len(elem) == 0 {
			// no match
			continue
		}

		// elem[0] is the entire line
		for i := 1; i < len(elem); i++ {
			fmt.Print(elem[i])
			fmt.Println()
		}
	}
	fmt.Println()
}
