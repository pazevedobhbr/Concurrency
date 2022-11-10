package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = "Hello, " + s

}

func printMessage() {
	fmt.Println(msg)
}

func main() {

	// challenge: modify this code so that the calls to updateMessage() on lines
	// 28, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	wg := sync.WaitGroup{}

	messages := []string{"universe!", "cosmos!", "world!"}
	for _, m := range messages {
		wg.Add(1)
		go updateMessage(m, &wg)
		wg.Wait()
		printMessage()
	}

}
