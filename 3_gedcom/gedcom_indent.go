package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
    "bytes"
	"strconv"
	"strings"
    "github.com/pborman/getopt"
)

// Thanks to this page for the original code:
// https://ebixio.com/blog/2012/03/05/editing-gedcom-files-with-vim
func main() {


    file := os.Stdin
    fi, err := file.Stat()
    if err != nil {
        fmt.Println("file.Stat()", err)
    }
    
    size := fi.Size()
    if size > 0 {
        fmt.Printf("%v bytes available in Stdin\n", size)
        
		in := bufio.NewReader(os.Stdin)
		for {
			s, err := in.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					log.Fatal(err)
				}
				break
			}
			trimmed := strings.TrimSpace(s)
			level, err := strconv.Atoi(strings.Split(trimmed, " ")[0])
			fmt.Print(strings.Repeat("\t", level))
			fmt.Print(trimmed)
			fmt.Print("\n")
		}
    } else {
        fmt.Println("Stdin is empty")
        
		//optName := getopt.StringLong("name", 'n', "Prakash", "Your name")
		//optPadding := getopt.StringLong("pad", 'p', "*", "Indentation")
		//home := os.UserHomeDir() 
		//file := getopt.StringLong("file", 'f', home + "/sarnobat.git/gedcom/rohidekar.ged", "Gedcom File")
		optHelp := getopt.BoolLong("help", 0, "Help")
		getopt.Parse()
		args := getopt.Args()

		if *optHelp {
			getopt.Usage()
			os.Exit(0)
		}

		//fmt.Println("File:\t" + *file)
		//fmt.Println("Name:\t" + *optName)
	
		// Get the remaining positional parameters
	
		fmt.Println("positional args: ", args)
		var fn string
		if (len(args) == 0) {
			home, _ := os.UserHomeDir() 
			fn = home + "/sarnobat.git/gedcom/rohidekar.ged"
		} else {
			fn = args[0]
		}
		
		
		file, err := os.Open(fn)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// Start reading from the file with a reader.
		reader := bufio.NewReader(file)
		for {
			var buffer bytes.Buffer

			var l []byte
			var isPrefix bool
			for {
				l, isPrefix, err = reader.ReadLine()
				buffer.Write(l)
				// If we've reached the end of the line, stop reading.
				if !isPrefix {
					break
				}
				// If we're at the EOF, break.
				if err != nil {
					if err != io.EOF {
						panic(err)
					}
					break
				}

			}
			line := buffer.String()


			// Process the line here.
//			fmt.Printf(" > Read %d characters\n", len(line))
//			fmt.Printf(" > > %s\n", limitLength(line, 50))

		
			
			if err == io.EOF {
				break
			}
			
			s:= line
			
			trimmed := strings.TrimSpace(s)
			level, _ := strconv.Atoi(strings.Split(trimmed, " ")[0])
			fmt.Print(strings.Repeat("\t", level))
			fmt.Print(trimmed)
			fmt.Print("\n")
			
		}
		if err != io.EOF {
			fmt.Printf(" > Failed with error: %v\n", err)
			panic(err)
		}

		
		
    }
    
}


func limitLength(s string, length int) string {
    if len(s) < length {
        return s
    }
    return s[:length]
}
