package main

import (
	"bufio"
	"io"
	"log"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	
	in := bufio.NewReader(os.Stdin)
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
		trimmed := strings.TrimSpace(s);
		level,err := strconv.Atoi(strings.Split(trimmed," ")[0])
        fmt.Print(strings.Repeat("\t",level))
        fmt.Print(trimmed)
        fmt.Print("\n")
	}


}

