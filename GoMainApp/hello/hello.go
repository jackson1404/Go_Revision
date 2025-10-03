package main

import (
	"fmt"
	"gomainapp/greetings"
	"log"
)

func main() {
	log.Printf("Reach %s", "here")
	message, err := greetings.Hello("Jackson")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
