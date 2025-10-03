package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary float32
}

type Employee struct {
	name   string
	salary float64
	age    int
}

func (e *Employee) Birthday() { // * refer to the memory address of original struct so it dont create copy of struct
	e.age += 1
}

func printPerson1Info(person1 Person) {
	fmt.Println("Name from method: ", person1.name)
	fmt.Println("Age: ", person1.age)
	fmt.Println("Job: ", person1.job)
	fmt.Println("Salary: ", person1.salary)
}

func main() {

	alice := &Employee{"w", 100000, 20} //Go automatically takes the address (&alice) for you.

	// alice := Employee{"w", 100000, 20}     with or without & is optional go will handle it

	alice.Birthday()
	fmt.Println("alice bir:", alice.age)

	var person1 Person
	var person2 Person

	// person1 specification
	person1.name = "Hege"
	person1.age = 45
	person1.job = "Teacher"
	person1.salary = 6000

	// person2 specification
	person2.name = "Cecilie"
	person2.age = 24
	person2.job = "Marketing"
	person2.salary = 4500

	printPerson1Info(person1)

	// Access and print person1 info
	fmt.Println("Name: ", person1.name)
	fmt.Println("Age: ", person1.age)
	fmt.Println("Job: ", person1.job)
	fmt.Println("Salary: ", person1.salary)

	// Access and print person2 info
	fmt.Println("Name: ", person2.name)
	fmt.Println("Age: ", person2.age)
	fmt.Println("Job: ", person2.job)
	fmt.Println("Salary: ", person2.salary)

	p := Person{"", 0, "", 0}
	fmt.Println("Person ja:", p)

	add := "22"         // value
	addMem := &add      // & store the memory address of value
	memValue := *addMem // * is dereference operator
	fmt.Println("adress mem:", addMem, memValue)

}
