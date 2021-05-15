package main
 
import (
	"bufio"
	"io"
	"log"
	"os"
	"fmt"
    "regexp"
	"github.com/jwangsadinata/go-multimap/slicemultimap"
)
 
func main() {
	in := bufio.NewReader(os.Stdin)
	
	mapp := slicemultimap.New()
	prevMd5 := "";
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
		r := regexp.MustCompile(`(?P<Md5>[^\s]+)\s+(?P<Path>.*)`)
		elem := r.FindStringSubmatch(s);
		
		//vals1, _ := mapp.Get(elem[1])		
		//fmt.Printf("Values: %d\n", len(vals1))
		countBefore := len(mapp.Values())
		//fmt.Println(mapp.Values());
		mapp.Put(elem[1], elem[2]);
		countAfter := len(mapp.Values())
		//vals2, _ := mapp.Get(elem[1])		
		//fmt.Printf("Values: %d\n", len(vals2))
		//fmt.Println(mapp.Values());
		
		if (prevMd5 == "") {
			fmt.Println("Initial - " + prevMd5)
		} else if (prevMd5 == elem[1]) {
			// Don't print anything, wait until the end
			//fmt.Println("Same, don't print - " + prevMd5)
			vals, _ := mapp.Get(elem[1])
			if (len(vals) > 1) {
			
				fmt.Printf("Values: %d\n", len(vals))
			}
			if (countBefore == 1 && countAfter == 1) {
				fmt.Println("error: ovewrote last value")
				//os.Exit(-1)
			}
		} else {
			prevMd5 = elem[1]
			
			// end of subsequence
			//paths, _ := mapp.Get(elem[1]);
			//if (len(paths) > 1) {
				fmt.Print ( countAfter);
				fmt.Print("\t");
				fmt.Print(elem[1]);
				fmt.Print("\t");
				fmt.Println(mapp.Values());
			
			//	fmt.Println("Not the same: " + prevMd5 + " -- " + elem[1])
			//}
			
			mapp.Clear()
		}
		//fmt.Println();		
		prevMd5 = elem[1]
		//fmt.Println("added: " + elem[1] + " :: " + elem[2])
	}
}
