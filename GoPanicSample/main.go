package main

import "fmt"

func divide(a, b int) int {
	if b == 0 {
		panic("division by zero") // 🚨 panic, like "throw" in Java
	}
	return a / b
}

func main() {
	defer func() {
		if r := recover(); r != nil { // like catch block in Java
			fmt.Println("Recovered from panic:", r)
		}
	}()

	result := divide(10, 0)
	fmt.Println("Result:", result) // won’t run unless recovered
}
