package main

import (
	"bufio"
	"fmt"
	"github.com/pborman/getopt"
	// 	"github.com/golang-collections/go-datastructures/queue"
	"io"
	"log"
	"os"
	"regexp"
)

func main() {
	optName := getopt.StringLong("name", 'n', "Sridhar", "Your name")
	optHelp := getopt.BoolLong("help", 0, "Help")
	getopt.Parse()

	if *optHelp {
		getopt.Usage()
		os.Exit(0)
	}
	fmt.Println("name = " + *optName)

	documentFrequenciesMap := make(map[string]int)

	//	var termFrequenciesChannel chan string = make(chan string)

	f, err := os.OpenFile("/tmp/queue.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	in := bufio.NewReader(os.Stdin)
	for {
		line, err := in.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		_ = line
		println("[debug] phase 1: " + line)

		regex := "^\\s*([0-9]+)*\\s*DOCUMENT_FREQUENCY_TOTAL..(.*)\n"
		r := regexp.MustCompile(regex)
		elem := r.FindStringSubmatch(line)

		if len(elem) == 0 {
			// no match

			// put it on the term frequencies channel
			// and delay processing (this will cause a queue backlog)

			//termFrequenciesChannel <- line
			if _, err := f.Write([]byte(line)); err != nil {
				log.Fatal(err)
			}

			continue
		} else {
			// process immediately, don't use a channel
			documentFrequenciesMap[elem[2]] += 1
		}
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	// Only read from the term frequencies channel when we've
	// finished processing the document frequencies channel

	// 	var c chan string = make(chan string)
	//
	// 	go func(c chan string) {
	// 		for i := 0; i < 3; i++ {
	// 			c <- "ping"
	//
	// 		}
	// 		close(c)
	// 	}(c)

	fmt.Println()
	fmt.Println("map:", documentFrequenciesMap)

	file, err := os.Open("/tmp/queue.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		println("[debug] phase 2: " + scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// 	for msg := range termFrequenciesChannel {
	// 		fmt.Println(msg)
	// 	}

	os.Exit(0)

}
