package main

import (
	"fmt"
	"log"

	"jackson.com/greetings"
)

func main() {
	message, err := greetings.Hello("Win")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(message)

}
