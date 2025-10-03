package main

import (
	"errors"
	"fmt"
	"log"
)

// 1. Normal function with parameters and return value
func add(a int, b int) int {
	return a + b
}

// 2. Function with multiple return values
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divided by zero")
	}
	return a / b, nil
}

// 3. Higher order function
func applyOperation(a int, b int, op func(int, int) int) int {
	return op(a, b)
}

// 4. filter example for passed fun
func filter(nunbers []int, test func(int) bool) []int {

	var result []int // init slice

	for _, v := range nunbers {
		if test(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {

	// call 1.
	sum := add(1, 2)
	fmt.Println("sum: ", sum)

	// call 2.
	divide, err := divide(10, 0)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Divide", divide)

	// Anonymous functions & Closures
	square := func(x int) int {
		return x * x
	}
	fmt.Println(square(4))

	// call 3.

	add := func(x, y int) int { return x + y }
	fmt.Println(add(3, 3))
	fmt.Println(applyOperation(8, 8, add))

	// num := []int{1, 2, 3, 4, 5}

	// for index, value := range num {
	// 	fmt.Printf("Index: %v - Value: %v \n", index, value)
	// }

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	even := filter(numbers, func(x int) bool { return x%2 == 0 })
	fmt.Println(even)

}
