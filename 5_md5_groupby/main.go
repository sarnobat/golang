package main

import (
	"bufio"
	"fmt"
	"github.com/jwangsadinata/go-multimap/slicemultimap"
	"io"
	"log"
	"os"
	"regexp"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	
	oneline := false;

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
		r := regexp.MustCompile(`(?P<Md5>[^\s]+)\s+(?P<Path>.*)`)
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
			
			if (oneline) {
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
