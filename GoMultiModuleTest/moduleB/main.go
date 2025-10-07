package main

import (
	"fmt"

	moduleA "jackson.com/moduleA"
)

func main() {
	fmt.Println("reach")
	msg, err := moduleA.Hello("jack")
	if err != nil {
		fmt.Println("argument not found:", err.Error())
		return
	}
	fmt.Println("Hello:", msg)
}
