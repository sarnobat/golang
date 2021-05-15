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

	mapp := slicemultimap.New()
	prevMd5 := ""
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

		r := regexp.MustCompile(`(?P<Md5>[^\s]+)\s+(?P<Path>.*)`)
		elem := r.FindStringSubmatch(s)

		countBefore := len(mapp.Values())
		mapp.Put(elem[1], elem[2])
		countAfter := len(mapp.Values())

		if prevMd5 == "" {
			fmt.Println("Initial - " + prevMd5)
		} else if prevMd5 == elem[1] {
			// Don't print anything, wait until the end
			vals, _ := mapp.Get(elem[1])
			if len(vals) > 1 {

				fmt.Printf("Values: %d\n", len(vals))
			}
			if countBefore == 1 && countAfter == 1 {
				fmt.Println("error: ovewrote last value")
				os.Exit(-1)
			}
		} else {
			prevMd5 = elem[1]

			// end of subsequence
			fmt.Print(countAfter)
			fmt.Print("\t")
			fmt.Print(elem[1])
			fmt.Print("\t")
			fmt.Println(mapp.Values())

			mapp.Clear()
		}
		prevMd5 = elem[1]
	}
}
