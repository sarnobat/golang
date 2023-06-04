//-----------------------------------------------------------------------------------------
// EXAMPLE
//
//	find ~/trash/ 	| go run ~/github/templates.git/helloworld.go
//	cat ~/.zshrc 	|  go run ~/github/templates.git/helloworld.go
//
// COMPILE TO NATIVE
//
//	# this will embed everything, as of 2020
//	go build helloworld.go
//
// DATE
//
//	2023 Jun 01
//

//-----------------------------------------------------------------------------------------

package main

import (
	"bufio"
	"fmt"
	// I forgot - why is the builtin "flag" package not good enougH?
	"github.com/pborman/getopt"
	"io"
	"log"
	"os"
	"strings"
)

var counts = make(map[string]int)

func main() {

	optHelp := getopt.BoolLong("help", 0, "Help")
	getopt.Parse()

	if *optHelp {
		getopt.Usage()
		os.Exit(0)
	}
	
	in := bufio.NewReader(os.Stdin)
	for {
		s, err := in.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		p := strings.TrimSpace(s)


		if _, err := os.Stat(p); !os.IsNotExist(err) {

			switch i, err := os.Stat(p); {
			case err != nil:
				fmt.Println(err)
			case i.IsDir():
			default:
				fmt.Println(p)
			}
		}
	}
}
