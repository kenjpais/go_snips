package main

import (
	"fmt"
	"strings"
)

// ping <-chan string notation for receive-only channel
// pong chan<- string notation for send-only channel

func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping
		pong <- strings.ToUpper(s)
	}
}

func main() {

	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something or press ENTER (enter Q to quit) ")
	for {
		fmt.Print("-> ")

		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput

		//wait for response
		response := <-pong
		fmt.Println("Response: ", response)
	}

	fmt.Println("All done. Closing channels.")
	close(ping)
	close(pong)
}
