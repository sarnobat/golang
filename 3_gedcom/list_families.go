package main

import (
	"fmt"
	"strings"
	"github.com/elliotchance/gedcom"
	"os"
    "github.com/pborman/getopt"
)

func main() {
    optName := getopt.StringLong("name", 'n', "Prakash", "Your name")
	optPadding := getopt.StringLong("pad", 'p', "  ", "Indentation")
	file := getopt.StringLong("file", 'f', "/Users/sarnobat/sarnobat.git/gedcom/rohidekar.ged", "Gedcom File")
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


	document, err := gedcom.NewDocumentFromGEDCOMFile("/Users/sarnobat/sarnobat.git/gedcom/rohidekar.ged")
	if err != nil {
		panic(err)
	}
	
	var root gedcom.FamilyNode
	individualFamilyMap  := make(map[string]gedcom.FamilyNode)


	for _, family := range document.Families() {
		fmt.Printf("%s\n", family)
		fmt.Printf("unique identifier: %s\n\n", family.Husband().Individual().UniqueIdentifiers().Strings()[0])
		if (family.Husband().Individual().UniqueIdentifiers().Strings()[0] == "d24b0564-1d9a-49b8-92de-9c381bf659a9") {
			root = *family;
			fmt.Printf("Found Root\n")
		}
		
		individualFamilyMap[family.Husband().Individual().String()] = *family
	}
	
	if (&root == nil) {
		panic("")
	}
	
	append(root, *optPadding, 1, individualFamilyMap)	
}

func append(familyNode gedcom.FamilyNode, indentation string, level int, individualFamilyMap map[string]gedcom.FamilyNode) {

	if (&familyNode != nil) {
		padding := strings.Repeat(indentation, level)
		for _, child := range familyNode.Children() {

			fmt.Printf("%s%s\n", padding, child)
			if val, ok := individualFamilyMap[child.Individual().String()]; ok {
				append(val, indentation, level + 1, individualFamilyMap)
			}
		}
	}	
}

