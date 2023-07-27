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
	"io"
	"log"
	"os"
	"strings"
	"strconv"
)

var counts = make(map[string]int)

func main() {
	
	if len(os.Args) < 1 {
		fmt.Println("Usage: ", os.Args[0], "arg1 arg2 ...")
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
		fmt.Println("Argument", i, ":", arg)
		if strings.HasPrefix(arg, "+") {
// 			n, err := strconv.Atoi(arg)
			lower1, err := humanSizeToBytes(removeFirstChar(arg))
			if err == nil {
				fmt.Println("lower bound: ", lower1)
				lower = lower1
			} else {
				log.Fatal(err)
				fmt.Println("Error 1")
				return;
			}
		} else if strings.HasPrefix(arg, "-") {

			upper1, err := humanSizeToBytes(removeFirstChar(arg))
			if err == nil {
				fmt.Println("upper bound: ", upper1)
				upper = upper1
			} else {
				log.Fatal(err)
				fmt.Println("Error 2")
				return;
			}
		} else {
			log.Fatal(err)
			fmt.Println("Error 3")
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
				fmt.Println(err)
			case fileInfo.Size() > lower && fileInfo.Size() < upper:
				fmt.Printf("in range: [%d] %d [%d]  %10s \n", lower, fileInfo.Size(), upper, p)
			case fileInfo.Size() > 10*1024*1024:
// 				fmt.Printf("File %s is bigger than 10MB (%d bytes)\n", filePath, fileInfo.Size())

			default:
				fmt.Printf("not in range: [%d] %d [%d]  %s\n",  lower, fileInfo.Size(), upper, p)
			}
		}
	}
}

func humanSizeToBytes(sizeStr string) (int64, error) {
	suffixes := map[string]int64{
		"":  1,
		"k": 1024,
		"M": 1024 * 1024,
		"G": 1024 * 1024 * 1024,
		"T": 1024 * 1024 * 1024 * 1024,
	}

	sizeStr = strings.ToUpper(sizeStr)

	for suffix, multiplier := range suffixes {
		if len(suffix) == 0 {
			sizeNum, _ := strconv.ParseInt(sizeStr, 10, 64)
			return sizeNum * multiplier, nil
		} else if len(suffix) > 0 && strings.HasSuffix(sizeStr, suffix) {
			fmt.Println("Before removing suffix ", suffix, " ", sizeStr)
			sizeNumStr := strings.TrimSuffix(sizeStr, suffix)
			fmt.Println("After suffix ", suffix, " ", sizeNumStr)
			sizeNum, err := strconv.ParseInt(sizeNumStr, 10, 64)
			if err != nil {
				fmt.Println("[error] 1 invalid size format: ", sizeNumStr, err)
				return 0, fmt.Errorf("[error] %s %s", err, sizeNumStr)
			}
			return sizeNum * multiplier, nil
		}
	}

	fmt.Println("[error] 2 invalid size format: %s", sizeStr)
	return 0, fmt.Errorf("invalid size format: %s", sizeStr)
}


func removeFirstChar(s string) string {
	if len(s) == 0 {
		return ""
	}
	return s[1:]
}
