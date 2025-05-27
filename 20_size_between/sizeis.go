//-----------------------------------------------------------------------------------------
// DESCRIPTION
//      Prints the path if the file is between the size intervals
//
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
	"io"
	"log"
	"os"
	"strings"
	"strconv"
	  "regexp"

)

var counts = make(map[string]int)

func main() {
	
	if len(os.Args) < 1 {
		fmt.Fprintln(os.Stderr, "Usage: ", os.Args[0], "<lower> <upper> ...")
		return
	}


	upper, err := humanSizeToBytes("999G")
	if err != nil {
		log.Fatal(err)
		return
	}
	lower, err := humanSizeToBytes("1")
	if err != nil {
		log.Fatal(err)
		return
	}
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
// 		fmt.Fprintln(os.Stderr, "Argument", i, ":", arg)
		if strings.HasPrefix(arg, "+") {
// 			n, err := strconv.Atoi(arg)
			lower1, err := humanSizeToBytes(removeFirstChar(arg))
			if err == nil {
// 				fmt.Fprintln(os.Stderr, "lower bound: ", lower1)
				lower = lower1
			} else {
				log.Fatal(err)
// 				fmt.Fprintln(os.Stderr, "Error 1")
				return;
			}
		} else if strings.HasPrefix(arg, "-") {

			upper1, err := humanSizeToBytes(removeFirstChar(arg))
			if err == nil {
// 				fmt.Fprintln(os.Stderr, "upper bound: ", upper1)
				upper = upper1
			} else {
				log.Fatal(err)
// 				fmt.Fprintln(os.Stderr, "Error 2")
				return;
			}
		} else {
// 			log.Fatal("Size arg needs + or -")
			fmt.Fprintln(os.Stderr, "Size arg needs + or -")
			return
		}
	}
	
	in := bufio.NewReader(os.Stdin)
	for {
		s, err := in.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		p := strings.TrimSpace(s)

		if _, err := os.Stat(p); !os.IsNotExist(err) {

			switch fileInfo, err := os.Stat(p); {
			case err != nil:
				fmt.Fprintln(os.Stderr, err)
			case fileInfo.Size() > lower && fileInfo.Size() < upper:
// 				fmt.Fprintf(os.Stderr, "in range: [%d] %d [%d]  %10s \n", lower, fileInfo.Size(), upper, p)
				fmt.Println(p)
			case fileInfo.Size() > 10*1024*1024:
// 				fmt.Printf("File %s is bigger than 10MB (%d bytes)\n", p, fileInfo.Size())

			default:
// 				fmt.Fprintf(os.Stderr, "not in range: [%d] %d [%d]  %s\n",  lower, fileInfo.Size(), upper, p)
			}
		}
	}
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
	
	
	
		if regexp.MustCompile(`\d$`).MatchString(sizeStr) {
			sizeNum, _ := strconv.ParseInt(sizeStr, 10, 64)
// 			fmt.Fprintln(os.Stderr, "[info] sizeNum = ", sizeNum)
			return sizeNum * multiplier, nil
		} else if len(suffix) > 0 && strings.HasSuffix(sizeStr, suffix) {
// 			fmt.Fprintln(os.Stderr, "Before removing suffix ", suffix, " ", sizeStr)
			sizeNumStr := strings.TrimSuffix(sizeStr, suffix)
// 			fmt.Fprintln(os.Stderr, "After suffix ", suffix, " ", sizeNumStr)
			sizeNum, err := strconv.ParseInt(sizeNumStr, 10, 64)
			if err != nil {
// 				fmt.Fprintln(os.Stderr, "[error] 1 invalid size format: ", sizeNumStr, err)
				return 0, fmt.Errorf("[error] %s %s", err, sizeNumStr)
			}			
// 			fmt.Fprintln(os.Stderr, "[info] sizeNum = ", sizeNum)
			return sizeNum * multiplier, nil
		}
	}

// 	fmt.Fprintln(os.Stderr, "[error] 2 invalid size format: ", sizeStr)
	return 0, fmt.Errorf("invalid size format: %s", sizeStr)
}


func removeFirstChar(s string) string {
	if len(s) == 0 {
		return ""
	}
	return s[1:]
}
