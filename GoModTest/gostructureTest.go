package main

import (
	"fmt"

	"rsc.io/quote/v4"
)

func main() {

	// Import external package and test

	fmt.Println(quote.Hello())

	// 1. Variable and Data types

	var age int = 20          // explicit type define (int)
	var degree float64 = 3.14 // (float)
	var isMale bool = true    // (boolean)
	var name = "Win"          // type inferred
	city := "YG"              // short declaration (work only inside method)
	var fileSize byte = 255
	var uniValue rune = 'A'

	fmt.Println("Name: " + name)
	fmt.Println("Age: ", age)
	fmt.Println("Degree: ", degree)
	fmt.Println("Is Male: ", isMale)
	fmt.Println("City: " + city)
	fmt.Println("Uni Value: ", uniValue)
	fmt.Println("file size: ", fileSize)

	// ===========================================

	// 2.Constants

	const Pi = 3.141
	const UserType = "GROUP"

	// Pi = Pi + 33   (Not allowed cuz const cant be changed value again once defined)

	fmt.Println("PI: ", Pi)
	fmt.Println("User type: ", UserType)

	// ===========================================

	// 3. Basic Operators

	a := 10
	b := 3

	// Arithmetic
	fmt.Println("Add: ", a+b)
	fmt.Println("Subtract: ", a-b)
	fmt.Println("Multiply:", a*b)
	fmt.Println("Divide:", a/b)
	fmt.Println("Modulus:", a%b)

	//Comparison
	fmt.Println("a==b? ", a == b)
	fmt.Println("a > b?", a > b)

	// Logical

	fmt.Println("a > 5 && b < 5? ", a > 5 && b < 5)
	fmt.Println("a < 5 || b > 5?", a < 5 || b < 5) // output true (One side to be true will return true)

	// ===========================================

	// 4. Conditional Statements (if/else)
	// Parenthese around the condition are optional.

	ageToCondition := 18

	if ageToCondition >= 18 {
		fmt.Println("you are adult")
	} else {
		fmt.Println("You are a minor")
	}

	// ===========================================

	// 5. Loops (for)
	// Go only has one looping construct: "for"

	for i := 1; i <= 5; i++ {
		fmt.Println("i=", i)
	}

	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println("sum= ", sum)

	// infinite loop
	// for {
	// 	fmt.Println("Looping forever")
	// }

	// ===========================================

	// 6. Arrays and Slices
	// Arrays -> fixed size, cannot grow
	// Slice -> Dynamic, can grow using append()

	// Arrays (fixed size)

	var numArray [3]int = [3]int{1, 2, 3}
	fmt.Println("arrays: ", numArray)

	// Arrays in Go are value types, which mean if you assign one array to another, it copies the data, doesnt reference the same memory

	var num2Array [3]int = numArray
	num2Array[0] = 100
	fmt.Println("Original Array: ", numArray)
	fmt.Println("Num 2 Array: ", num2Array)

	// Arrays Access
	// index start at 0
	// len() gives the size of array

	fmt.Println("First element: ", numArray[0])
	fmt.Println("Length of Array:", len(numArray))

	// Slice (Dynamic size)

	fruits := []string{"Apple", "Banana", "Cherry"} // Declares a slice of strings array
	fruits = append(fruits, "Mango")                // Add "Mango" to the slice as slices are dynamic can grow or shrink.
	fmt.Println("Slice: ", fruits)

	// Slices are reference types, meaning if you copy a slice, both the value and reference point to the same underlying arrary
	moreFruits := fruits
	fmt.Println("Origin fruite before add: ", fruits)
	moreFruits[0] = "WaterMelon"
	fmt.Println("Added fruit list: ", fruits)
	fmt.Println("More fruit list", moreFruits)

	// Slice Properties

	fmt.Println("Length: ", len(fruits))
	fmt.Println("Capacity:", cap(fruits)) // check how much size this slice can still hold (in total) before Go has to allocate new memory

}
