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
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		// TODO: change this to support delimiter "::" so that we can
		// use this on exif_gps output and group photos (then display them on google maps)
		r := regexp.MustCompile(`(?P<Md5>[^\s]+)\s+(?P<Path>.*)`)
		elem := r.FindStringSubmatch(s)

		countBefore := len(mapp.Values())
		mapp.Put(elem[1], elem[2])
		countAfter := len(mapp.Values())

		if prevMd5 == "" {
			// Don't print anything before seeing the 2nd row
		} else if prevMd5 == elem[1] {
			// Don't print anything, wait until the end
			if countBefore == 1 && countAfter == 1 {
				fmt.Println("error: ovewrote last value")
				os.Exit(-1)
			}
		} else {
			// Print the aggregate, end of subsequence
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
