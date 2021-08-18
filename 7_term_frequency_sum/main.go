package main

import (
	"bufio"
	"fmt"
	"github.com/pborman/getopt"
	"io"
	"log"
	"os"
	"regexp"
)

func pinger(c chan string) {
	for i := 0; i < 3; i++ {
		c <- "ping"
	}
	close(c)
}

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
	//     messages := make(chan string)
	//
	// 	go func() {
	// 	    msg := <-messages
	// 		fmt.Println(msg)
	// 	}()
	//
	// 	messages <- "ping"

	var c chan string = make(chan string)

	go func(c chan string) {
		for i := 0; i < 3; i++ {
			c <- "ping"
		}
		close(c)
	}(c)

	for msg := range c {
		fmt.Println(msg)
	}

	os.Exit(0)

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
		println("[debug] line: " + line)

		regex := "^\\s*([0-9]+)*\\s*DOCUMENT_FREQUENCY_TOTAL..(.*)\n"
		r := regexp.MustCompile(regex)
		elem := r.FindStringSubmatch(line)

		if len(elem) == 0 {
			// no match
			// 			messages <- "foo"
			continue
		} else {

			// 			messages <- "bar"

			documentFrequenciesMap[elem[2]] += 1
			// elem[0] is the entire line
			for i := 1; i < len(elem); i++ {
				fmt.Print(elem[i])
				fmt.Println()
			}
		}
	}

	fmt.Println()
	fmt.Println("map:", documentFrequenciesMap)

}
