package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s, ok := <-ping
		if !ok {
			fmt.Println("Ping channel closed. Exiting...")
			return
		}
		pong <- fmt.Sprintf("%s!", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type Something and ENTER (enter Q to quit)")
	for {
		// print a prompt
		fmt.Print("-> ")

		// get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)
		if userInput == strings.ToLower("q") {
			break
		}
		ping <- userInput
		//wait a response
		response := <-pong
		fmt.Println("Response:", response)
	}
	fmt.Println("All done! Closing channels...")
	close(ping)
	close(pong)
}
