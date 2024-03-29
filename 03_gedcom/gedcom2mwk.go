package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/pborman/getopt"
	"os"
	"strings"
)

func main() {
	optName := getopt.StringLong("name", 'n', "Prakash", "Your name")
	optPadding := getopt.StringLong("pad", 'p', "*", "Indentation")
	home, _ := os.UserHomeDir()
	file := getopt.StringLong("file", 'f', home+"/sarnobat.git/gedcom/rohidekar.ged", "Gedcom File")
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

	document, err := gedcom.NewDocumentFromGEDCOMFile(*file)
	if err != nil {
		panic(err)
	}

	var root gedcom.FamilyNode
	var rootIndividual gedcom.IndividualNode
	individualFamilyMap := make(map[string]gedcom.FamilyNode)

	for _, individual := range document.Individuals() {
		fmt.Printf("individual: %s\n", individual.String())
	}

	for _, family := range document.Families() {
		fmt.Printf("%s\n", family)
		fmt.Printf("unique identifier: %s\n\n", family.Husband().Individual().UniqueIdentifiers().Strings()[0])
		if family.Husband().Individual().UniqueIdentifiers().Strings()[0] == "799db437-e0d2-44cc-a8f9-afda533cb5b7" {
			root = *family
			rootIndividual = *family.Husband().Individual()
			fmt.Printf("Found Root\n")
		}

		//		individualFamilyMap[family.Husband().Individual().String()] = *family
		//		individualFamilyMap[family.Wife().Individual().String()] = *family
		individualFamilyMap[family.Husband().Individual().UniqueIdentifiers().Strings()[0]] = *family
		individualFamilyMap[family.Wife().Individual().UniqueIdentifiers().Strings()[0]] = *family
		fmt.Println("[DEBUG] " + family.Husband().Individual().String() + " :: " + family.String())
		fmt.Println("[DEBUG] " + family.Wife().Individual().String() + " :: " + family.String())
	}

	if &root == nil {
		panic("")
	}
	fmt.Printf("\n")
	append(root, *optPadding, 1, individualFamilyMap)

	printIndividual(rootIndividual, *optPadding, 1, individualFamilyMap)
}

func printIndividual(rootIndividual gedcom.IndividualNode, indentation string, level int, individualFamilyMap map[string]gedcom.FamilyNode) {

	padding := strings.Repeat(indentation, level)
	fmt.Printf("%s %s", padding, rootIndividual.String())

	familyNode, found := individualFamilyMap[rootIndividual.UniqueIdentifiers().Strings()[0]]
	if found {
		// TODO: iterate over all spouses
		spouse := *rootIndividual.Spouses()[0]
		fmt.Printf("  (-- %s)", spouse.Name().String())
		fmt.Printf("\n")

		if &familyNode != nil {
			for _, child := range familyNode.Children() {
				printIndividual(*child.Individual(), indentation, level+1, individualFamilyMap)
			}
		}
	} else {
		fmt.Printf("\n")
	}
}

func append(familyNode gedcom.FamilyNode, indentation string, level int, individualFamilyMap map[string]gedcom.FamilyNode) {

	if &familyNode != nil {

		padding := strings.Repeat(indentation, level)
		for _, child := range familyNode.Children() {

			fmt.Printf("%s %s\n", padding, child)
			if val, ok := individualFamilyMap[child.Individual().String()]; ok {
				append(val, indentation, level+1, individualFamilyMap)
			}
		}
	}
}
