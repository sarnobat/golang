package main

import (
	"bufio"
	"fmt"
	"github.com/pborman/getopt"
	"io"
	"log"
	"os"
	"regexp"
)

func main() {

	optDelimiter := *getopt.StringLong("pad", 'd', "\\*", "Character sequence to be expanded to parent text")
	optHelp := getopt.BoolLong("help", 0, "Help")
	getopt.Parse()

	if *optHelp {
		getopt.Usage()
		os.Exit(0)
	}

	in := bufio.NewReader(os.Stdin)
	pathSegments := make([]string, 9)

	for {
		s, err := in.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		exp := "^(" + optDelimiter + "*)\\s*(.*)"
		r := regexp.MustCompile(exp)
		elem := r.FindStringSubmatch(s)

//		fmt.Println()

		if len(elem) == 0 {
			continue
		}
		fmt.Println(len(elem[1]))
		
		pathSegments[len(elem[1])] = elem[2];
		
		for i := 0; i < len(elem[1]); i++ {
			fmt.Print(pathSegments[i])
			fmt.Print("/")
		}
		fmt.Print(elem[2])
	}
}
