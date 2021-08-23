package main

import (
    "strconv"

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
		print("[debug] phase 1: " + line)

		regex := "^\\s*([0-9]+)*\\s*DOCUMENT_FREQUENCY_TOTAL..(.*)\n"
		r := regexp.MustCompile(regex)
		elem := r.FindStringSubmatch(line)

		if len(elem) == 0 {
			// not a document frequency row, is a term frequency row
			// just write it out to a file and scan through it later
			if _, err := f.Write([]byte(line)); err != nil {
				log.Fatal(err)
			}
		} else {
			documentFrequenciesMap[elem[2]] += 1
		}
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("map:", len(documentFrequenciesMap))

	file, err := os.Open("/tmp/queue.log")
	if err != nil {
		log.Fatal(err)
	}
	err = file.Truncate(0)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
// 		println("[debug] phase 2: line = " + line)
		
		//regex := "^\\s*([0-9]+)*\\s*(.*): (.*)\n"
		regex := "^\\s*([0-9]+)\\s+(.*):\\s+(.*)"
		r := regexp.MustCompile(regex)
		elem := r.FindStringSubmatch(line)

		if len(elem) == 0 {
			println("[debug] phase 2: skipping line = ", line)
// 			println("[debug] phase 2: term frequency = ", elem[1])
			os.Exit(-1)
// 			continue
		} else {

// 			println("[debug] phase 2: total docs: ", len(documentFrequenciesMap))
// 			println("[debug] phase 2: term frequency = ", elem[1])
// 			println("[debug] phase 2: document = ", elem[2])
// 			println("[debug] phase 2: term = ", elem[3])
// 			println("[debug] phase 2: document frequency = ", documentFrequenciesMap[elem[3]])
// 			
			tf, err := strconv.ParseFloat(elem[1],8)
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}

			df := documentFrequenciesMap[elem[3]]
			
			if (df == 0) {
				fmt.Println(err)
				fmt.Println("Couldn't find document frequency, skipping: " + line)
				fmt.Println("Elems: ", elem)
				fmt.Printf("Elem 3: >%s<", elem[3])
				os.Exit(2)
// 				continue
			}

			// (what does df = 0 mean? We are getting this case)
			// TODO: why is everything a round number in the output?
			score := float64(tf) / float64(df)
// 			fmt.Printf("phase 2: tfidf score = %f\n", score)
			fmt.Printf("phase 3: %s\t%.1f\t%s\n", elem[2], score, elem[3])

// 			println("phase 2: tfidf score = ", score)
			println()
// 			os.Exit(-1)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
