package main

import (
	"fmt"

	"jackson.com/greetings"
)

func main() {
	message := greetings.Hello("Jackson")
	fmt.Println(message)
}
