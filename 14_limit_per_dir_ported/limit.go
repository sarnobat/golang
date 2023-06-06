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
//	2022 Jan 20
//

//-----------------------------------------------------------------------------------------

package main

import (
	"bufio"
	"fmt"
	// I forgot - why is the builtin "flag" package not good enougH?
	"github.com/pborman/getopt"
// 	"io"
// 	"log"
	"os"
// 	"os/exec"
	"path/filepath"
// 	"regexp"
// 	"strings"
)

var counts = make(map[string]int)

// golang MUST have a main function (unlike python)
func main() {

	///
	/// 5) CLI options
	///
	optName := getopt.StringLong("name", 'n', "", "Your name")
	optHelp := getopt.BoolLong("help", 0, "Help")
	getopt.Parse()

	if *optHelp {
		getopt.Usage()
		os.Exit(0)
	}
	fmt.Println("Hello " + *optName)

    // Initialize the variables
    numMax := 5
    filesInPath := make(map[string]int)
    lastPath := ""
//     count := 0

    // Read the lines from the standard input
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        // Get the path
        path := filepath.Dir(scanner.Text())

        // Check if the path is the same as the previous path
        if lastPath == path {
            // The path is the same, increment the counter
            filesInPath[path]++
        } else {
            // The path is different, reset the counter
            filesInPath[path] = 1
            lastPath = path
        }

        // Check if the counter is greater than the maximum number
        if filesInPath[path] > numMax {
            // The counter is greater than the maximum number, continue to the next line
            continue
        }

        // Print the line
        fmt.Println(scanner.Text())
    }

// 	///
// 	/// 1) Loop over stdin
// 	///
// 	in := bufio.NewReader(os.Stdin)
// 	for {
// 		s, err := in.ReadString('\n')
// 		if err != nil {
// 			// io.EOF is expected, anything else
// 			// should be handled/reported
// 			if err != io.EOF {
// 				log.Fatal(err)
// 			}
// 			break
// 		}
// 		// Do something with the line of text
// 		// in string variable s.
// 		_ = s
// 		p := strings.TrimSpace(s)
// 
// 		//
// 		// 3) Parse file path
// 		//
// 
// 		if file, err := os.Stat(p); !os.IsNotExist(err) {
// 
// 			// check if it's a file or a dir
// 			switch i, err := os.Stat(p); {
// 			case err != nil:
// 				fmt.Println(err)
// 			case i.IsDir():
// 				fmt.Println("	", p, "is a directory")
// 			default:
// 				fmt.Println("	", p, "is a file")
// 			}
// 			// get absolute path
// 			abs, _ := filepath.Abs(p)
// 			fmt.Println("	absolute: " + abs)
// 
// 			///
// 			/// 6) Call a shell program instead
// 			///
// 			cmd := exec.Command("dirname", p)
// 			out, err := cmd.CombinedOutput()
// 			if err != nil {
// 				log.Fatal(err, out)
// 			}
// 
// 			cmd2 := exec.Command("basename", p)
// 			out2, err2 := cmd2.CombinedOutput()
// 			if err2 != nil {
// 				log.Fatal(err2)
// 			}
// 			dirname := strings.Trim(string(out2), "\n")
// 			//			fmt.Println("	DIRNAME: ", dirname)
// 
// 			///
// 			/// 5) dictionary
// 			///
// 			_, exist := counts[dirname]
// 			if !exist {
// 				counts[dirname] = 1
// 			} else {
// 				counts[dirname]++
// 			}
// 
// 			// We don't need to check if it's a folder to recurse into. Stdin will ensure that
// 			// we recurse.
// 			if counts[dirname] <= 5 {
// 				fmt.Println("", p, "\n")
// 			}
// 
// 		} else {
// 			///
// 			/// 1) Print to stdout
// 			///
// 			if file == nil {
// 				fmt.Print("from go: ", p, "\n")
// 			}
// 
// 			///
// 			/// 2) Regex capture groups extracted and read separately (not used directly
// 			/// in a substitution) will cover more scenarios.
// 			///
// 			regex := "^\\s*([0-9]+)*\\s*DOCUMENT_FREQUENCY_TOTAL..(.*)\n"
// 			r := regexp.MustCompile(regex)
// 			line := s
// 			elem := r.FindStringSubmatch(line)
// 
// 			if len(elem) == 0 {
// 				// no match
// 				continue
// 			}
// 
// 			// elem[0] is the entire line
// 			for i := 1; i < len(elem); i++ {
// 				fmt.Print(elem[i])
// 				fmt.Println()
// 			}
// 		}
// 	}
}
