//-----------------------------------------------------------------------------------------
// EXAMPLE
//
//	find ~/trash/ 	| go run ~/github/templates.git/helloworld.go
//	cat ~/.zshrc 	|  go run ~/github/templates.git/helloworld.go
//
// COMPILE TO NATIVE
//
//	# this will embed everything, as of 2020
//	go build helloworld.go
//
// DATE
//
//	2023 Jun 01
//
//-----------------------------------------------------------------------------------------

package main

import (
	"bufio"
	"fmt"
	// I forgot - why is the builtin "flag" package not good enougH?
	"github.com/pborman/getopt"
	"io"
	"log"
	"os"
	"strings"
	"strconv"
)

var counts = make(map[string]int)

func main() {

	optHelp := getopt.BoolLong("help", 0, "Help")
	getopt.Parse()

	if *optHelp {
		getopt.Usage()
		os.Exit(0)
	}
	
	if len(os.Args) < 1 {
		fmt.Println("Usage: ", os.Args[0], "arg1 arg2 ...")
		return
	}


	upper, _ := humanSizeToBytes("100G")

}

func humanSizeToBytes(sizeStr string) (int64, error) {
	suffixes := map[string]int64{
		"":  1,
		"K": 1024,
		"M": 1024 * 1024,
		"G": 1024 * 1024 * 1024,
		"T": 1024 * 1024 * 1024 * 1024,
	}

	sizeStr = strings.ToUpper(sizeStr)

	for suffix, multiplier := range suffixes {
		if strings.HasSuffix(sizeStr, suffix) {
			sizeNumStr := strings.TrimSuffix(sizeStr, suffix)
			sizeNum, err := strconv.ParseInt(sizeNumStr, 10, 64)
			if err != nil {
				return 0, err
			}
			return sizeNum * multiplier, nil
		}
	}

	return 0, fmt.Errorf("invalid size format: %s", sizeStr)
}

