package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, ":", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// Run printMessage concurrently
	go printMessage("goroutine")

	// Run printMessage in main thread
	printMessage("main function")

	// Give goroutine time to finish before main exits
	time.Sleep(3 * time.Second)
}
