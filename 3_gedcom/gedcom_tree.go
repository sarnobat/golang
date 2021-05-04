package main

import (
	/*	"bufio"
		"io"
		"log"
		"os"*/
	"encoding/json"
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/pborman/getopt"
	"os"
	"reflect"
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

	fmt.Println("File:\t" + *file)
	fmt.Println("Name:\t" + *optName)

	// Get the remaining positional parameters

	fmt.Println("positional args: ", args)

	document, err := gedcom.NewDocumentFromGEDCOMFile("/Users/sarnobat/sarnobat.git/gedcom/rohidekar.ged")
	if err != nil {
		panic(err)
	}
	//reflect.TypeOf(document)
	fmt.Print(reflect.TypeOf(document))
	//fmt.Print(document)

	for _, individual := range document.Individuals() {
		//		data,_ := json.Marshal(individual)
		//		fmt.Printf("%s\n", data)

		fmt.Println(individual.UniqueIdentifiers())
		fmt.Println(reflect.TypeOf(individual.UniqueIdentifiers()))
		fmt.Println(*individual.UniqueIdentifiers())
		fmt.Println((*individual.UniqueIdentifiers()).Strings())
		fmt.Println((*individual.UniqueIdentifiers()).String())

		gett := document.Individuals().ByUniqueIdentifier((*individual.UniqueIdentifiers()).String())
		//		gett :=document.Individuals().ByUniqueIdentifier("(603c08de-ccff-91eb-2389-000000000000)");
		//				gett :=document.Individuals().ByUniqueIdentifier("");

		//		fmt.Println(gett);
		data1, _ := json.Marshal(gett)
		fmt.Printf("%s\n", data1)

		fmt.Println()

		//fmt.Println(individual.UniqueIDs())
		//fmt.Printf("%#v\n", individual)

	}
	/*
		str := "";
		for _, fam := range document.Families() {
		  str += getFamilyAsString(fam);
		}*/
	//	print(str)
	//	a :=document.Individuals().ByUniqueIdentifier("(603c08de-ccff-91eb-2389-000000000000)");
	// 	print(a);
	// 	for _, fam := range document.Families() {
	//	}

}

/*
func getFamilyAsString(family *gedcom.FamilyNode) string {
	print((*family).String());
	ret := family.String() + "\n";
	for _,child := range family.Children() {
		if (child.Family() == family) {
			//os.Exit(-1);
		}
		//ret = ret + getFamilyAsString(child.Individual().());
	}
	return ret;

}
*/
