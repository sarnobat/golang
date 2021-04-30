package main

import (
/*	"bufio"
	"io"
	"log"
	"os"*/
	"fmt"
	"github.com/elliotchance/gedcom"
)

func main() {
	document, err := gedcom.NewDocumentFromGEDCOMFile("family.ged")
	if err != nil {
		panic(err)
	}
	fmt.Print(document)

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

