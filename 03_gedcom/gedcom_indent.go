package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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
    var in *bufio.Reader
    
    size := fi.Size()
    if size > 0 {
        fmt.Fprintf(os.Stderr, "[DEBUG] %v bytes available in Stdin\n", size)
        
		in = bufio.NewReader(os.Stdin)
		
    } else {
        fmt.Fprintf(os.Stderr, "[DEBUG] Stdin is empty\n")
        
		optHelp := getopt.BoolLong("help", 0, "Help")
		getopt.Parse()
		args := getopt.Args()

		if *optHelp {
			getopt.Usage()
			os.Exit(0)
		}	
		// Get the remaining positional parameters
	
		fmt.Fprintf(os.Stderr, "[DEBUG] positional args: %s\n", args)
		var fn string
		if (len(args) == 0) {
			home, _ := os.UserHomeDir() 
			fn = home + "/sarnobat.git/2023/genealogy/rohidekar.ged"
		} else {
			fn = args[0]
		}
		
		
		file, err := os.Open(fn)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		in = bufio.NewReader(file)
    }
        
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
}
