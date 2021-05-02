package main

import (
	"bufio"
	"io"
	"log"
	"fmt"
	"github.com/elliotchance/gedcom"
	"os"
    "github.com/pborman/getopt"
	"reflect"
	"strings"
	"strconv"
)

func main() {
    optName := getopt.StringLong("name", 'n', "Prakash", "Your name")
	file := getopt.StringLong("file", 'f', "/Users/sarnobat/sarnobat.git/gedcom/rohidekar.ged", "Gedcom File")
    optHelp := getopt.BoolLong("help", 0, "Help")
    getopt.Parse()
	args := getopt.Args()

    if *optHelp {
        getopt.Usage()
        os.Exit(0)
    }

    println("File:\t" + *file)
    println("Name:\t" + *optName)
    
    // Get the remaining positional parameters
	
	println("positional args: ", args)


	document, err := gedcom.NewDocumentFromGEDCOMFile("/Users/sarnobat/sarnobat.git/gedcom/rohidekar.ged")
	if err != nil {
		panic(err)
	}
	//reflect.TypeOf(document)
	print(reflect.TypeOf(document))
	//print(document)
	
	for _, individual := range document.Individuals() {
	  println(individual)
	  //println("")
	}
	
	
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
		//print("added: "+s)
		
		trimmed := strings.TrimSpace(s);
		level,err := strconv.Atoi(strings.Split(trimmed," ")[0])
        print(strings.Repeat("\t",level))
        print(trimmed)
        print("\n")
	}


}

