package main

import (
/*	"bufio"
	"io"
	"log"
	"os"*/
	"fmt"
// 	"github.com/elliotchance/gedcom"
	"os"
    "github.com/pborman/getopt"

)

func main() {
    optName := getopt.StringLong("name", 'n', "", "Your name")
	file := getopt.StringLong("file", 'f', "", "Gedcom File")
    optHelp := getopt.BoolLong("help", 0, "Help")
    getopt.Parse()
	args := getopt.Args()

    if *optHelp {
        getopt.Usage()
        os.Exit(0)
    }

    fmt.Println("File:\t" + *file)
    fmt.Println("Name:\t" + *optName)
    
    // Get the remaining positional parameters
	
	fmt.Println("positional args: ", args)


// 	document, err := gedcom.NewDocumentFromGEDCOMFile("family.ged")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Print(document)

/*
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
		fmt.Print("added: "+s)
	}
*/
}

