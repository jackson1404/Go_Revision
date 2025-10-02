package main

import (
	"fmt"
	"gomainapp/greetings"
)

func main() {
	message := greetings.Hello("Jackson")
	fmt.Println(message)
}
