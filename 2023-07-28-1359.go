package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello World!")

	// Exercise 2: Variables
	var name string = "John"
	age := 25
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// Exercise 3: Arithmetic Operations
	a := 10
	b := 5
	sum := a + b
	diff := a - b
	product := a * b
	quotient := a / b
	remainder := a % b
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", diff)
	fmt.Printf("Product: %d\n", product)
	fmt.Printf("Quotient: %d\n", quotient)
	fmt.Printf("Remainder: %d\n", remainder)

	// Exercise 4: Conditional Statements
	x := 5
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than or equal to 10")
	}

	// Exercise 5: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Exercise 6: Arrays
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	fmt.Println(numbers)

	// Exercise 7: Slices
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// Exercise 8: Maps
	phonebook := make(map[string]string)
	phonebook["John"] = "123-456-7890"
	phonebook["Jane"] = "987-654-3210"
	fmt.Println(phonebook)
}