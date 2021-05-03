package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"os"
)

func main() {

	document, err := gedcom.NewDocumentFromGEDCOMFile("rohidekar.ged")
	if err != nil {
		panic(err)
	}
	for _, individual := range document.Individuals() {
	
		i := document.Individuals().ByUniqueIdentifier((*individual.UniqueIdentifiers()).String());
		
		if (i == nil) {
			fmt.Println("Cannot get individual by ID: " + (*individual.UniqueIdentifiers()).String());
			os.Exit(-1)
		}		
	}
}
