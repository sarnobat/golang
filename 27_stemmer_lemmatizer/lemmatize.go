package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

	"github.com/aaaton/golem/v4"
	"github.com/aaaton/golem/v4/dicts/en"
)

func main() {

	// the language packages are available under golem/dicts
	// "en" is for english
	g, err := golem.New(en.New())
	if err != nil {
		panic(err)
	}
	
    // Initialize the golem lemmatizer
//     g, err := golem.New()
//     if err != nil {
//         fmt.Println("Error initializing golem:", err)
//         return
//     }

    // Read input from stdin
    scanner := bufio.NewScanner(os.Stdin)
//     fmt.Println("Please enter some text:")

    for scanner.Scan() {
        input := scanner.Text()

        // Tokenize the input into words
        words := strings.Fields(input)

        // Lemmatize each word
        for _, word := range words {
//             lemma := g.Lemmatize(word)
           	out := g.Lemma(word)

//            fmt.Printf("Original: %s -> Lemma: %s\n", word, out)
			fmt.Printf("%s\n", out)
			fmt.Fprintf(os.Stderr, "[debug] %s -> %s\n", word, out)
        }

//         break // Exit after processing one line (you can modify this for continuous input)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading from stdin:", err)
    }
}

