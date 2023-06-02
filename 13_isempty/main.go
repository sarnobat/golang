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
    "io/ioutil"
)

var counts = make(map[string]int)

// golang MUST have a main function (unlike python)
func main() {

	///
	/// 5) CLI options
	///
// 	optName := getopt.StringLong("invert", 'v', "", "Invert match")
	optHelp := getopt.BoolLong("help", 0, "Help")
	optInvert := getopt.BoolLong("invert", 'v', "Invert match")
	getopt.Parse()

	if *optHelp {
		getopt.Usage()
		os.Exit(0)
	}
	
	result := fmt.Sprintf("invert=%t", *optInvert)
	fmt.Println(result)

	///
	/// 1) Loop over stdin
	///
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
		p := strings.TrimSpace(s)

		//
		// 3) Parse file path
		//

		if file, err := os.Stat(p); !os.IsNotExist(err) {

			switch i, err := os.Stat(p); {
			case err != nil:
				fmt.Println(err)
			case i.IsDir():
				files, err := ioutil.ReadDir(p)
				if err != nil {
					fmt.Println(err)
					return
				}
				if len(files) == 0 {
					fmt.Fprintf(os.Stderr, "[empty dir]" + p + "\n")
					fmt.Println(p)
				} else {
					fmt.Fprintf(os.Stderr, "[not empty dir]" + p + "\n")
				}

			default:
				if file.Size() == 0 {
// 					fmt.Println(p)
					fmt.Fprintf(os.Stderr, "[empty file]" + p + "\n")
				} else {
					fmt.Fprintf(os.Stderr, "[not empty file] " + p + "\n")
				}
			}
		}
	}
}
