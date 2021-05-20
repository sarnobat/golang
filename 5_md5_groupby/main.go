package main

import (
	"bufio"
	"fmt"
	"github.com/jwangsadinata/go-multimap/slicemultimap"
	"github.com/pborman/getopt"
	"io"
	"log"
	"os"
	"regexp"
)

func main() {

	optDelimiter := "\\s+"
	optDelimiter2 := getopt.StringLong("delimiter", 'd', "\\s+", "Group file paths onto a single line")

	optHelp := getopt.BoolLong("help", 0, "Help")
	oneline := getopt.BoolLong("oneline", 0, "Group file paths onto a single line")
	getopt.Parse()

	if *optHelp {
		getopt.Usage()
		os.Exit(0)
	}

	if optDelimiter == *optDelimiter2 {
		optDelimiter = *optDelimiter2
		//	println("all good: optDelimiter = " + optDelimiter)
	} else {
		//	println("Using a different delimiter")
	}

	in := bufio.NewReader(os.Stdin)

	mapp := slicemultimap.New()
	prevMd5 := ""
	for {
		s, err := in.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		// TODO: change this to support delimiter "::" so that we can
		// use this on exif_gps output and group photos (then display them on google maps)
		//delim := "\s+"
		//exp := `(?P<Md5>[^\s]+)\s+(?P<Path>.*)`
		// 		optDelimiter2 := "\\s+"
		// 		fmt.Fprintf(os.Stderr, "[DEBUG] optDelimiter2 = %v\n", optDelimiter2)
		exp := "(?P<Md5>[^\\s]+)" + optDelimiter + "(?P<Path>.*)"
		fmt.Fprintf(os.Stderr, "[DEBUG] exp = %v\n", exp)
		r := regexp.MustCompile(exp)
		elem := r.FindStringSubmatch(s)

		// fmt.Fprintf(os.Stderr, "[DEBUG] elem[1] = %v\n", elem[1])
		// fmt.Fprintf(os.Stderr, "[DEBUG] elem[2] = %v\n", elem[2])
		// fmt.Fprintf(os.Stderr, "[DEBUG]\n")

		if prevMd5 == "" {
			// Don't print anything before seeing the 2nd row
		} else if prevMd5 == elem[1] {
			// fmt.Fprintf(os.Stderr, "[DEBUG] prevMd5 == elem[1] = %v\n", prevMd5,elem[1])
			// Don't print anything, wait until the end

		} else {
			// Print the aggregate, end of subsequence
			prevValues, _ := mapp.Get(prevMd5)

			if *oneline {
				fmt.Print(len(prevValues))
				fmt.Print("\t")
				fmt.Print(prevMd5)
				fmt.Print("\t")
				fmt.Println(prevValues)
			} else {
				fmt.Println()
				for _, s := range prevValues {

					fmt.Print(len(prevValues))
					fmt.Print("\t")
					fmt.Print(prevMd5)
					fmt.Print("\t")
					fmt.Println(s)
				}
				fmt.Print(len(prevValues))
				fmt.Print("\t")
				fmt.Print(prevMd5)
				fmt.Print("\t")
			}

			mapp.Clear()
		}
		countBefore := len(mapp.Values())
		// 		fmt.Fprintf(os.Stderr, "[DEBUG] mapp = %v\n", mapp)

		mapp.Put(elem[1], elem[2])
		// 		fmt.Fprintf(os.Stderr, "[DEBUG] mapp = %v\n", mapp)
		countAfter := len(mapp.Values())

		if countBefore == 1 && countAfter == 1 {
			fmt.Println("error: ovewrote last value")
			os.Exit(-1)
		}
		prevMd5 = elem[1]

	}
}
