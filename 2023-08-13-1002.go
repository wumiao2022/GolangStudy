package main

import "fmt"

func main() {
	fmt.Println("Hello, Gopher!")

	// Declare a variable of type string
	var name string

	// Assign a value to the variable
	name = "John Doe"

	// Print the value of the variable
	fmt.Println("My name is", name)

	// Declare and assign a variable in one line
	age := 30

	// Print the value of the variable
	fmt.Println("I am", age, "years old")

	// Declare multiple variables in one line
	var (
		height  float64 = 180.5
		weight  int     = 75
		isMale  bool    = true
		nationality string  = "British"
	)

	// Print the values of the variables
	fmt.Println("Height:", height)
	fmt.Println("Weight:", weight)
	fmt.Println("Is male:", isMale)
	fmt.Println("Nationality:", nationality)
}

// Output:
// Hello, Gopher!
// My name is John Doe
// I am 30 years old
// Height: 180.5
// Weight: 75
// Is male: true
// Nationality: British