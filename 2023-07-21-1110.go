package main

import "fmt"

func main() {
	// Exercise 1: Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Exercise 2: Declare variables and assign values
	var a int = 10
	var b float64 = 3.14
	var c bool = true
	var d string = "Golang"
	fmt.Println(a, b, c, d)

	// Exercise 3: Perform calculations
	sum := a + int(b)
	fmt.Println("The sum is:", sum)

	// Exercise 4: Use if-else condition
	if a > 5 {
		fmt.Println("Variable a is greater than 5")
	} else {
		fmt.Println("Variable a is less than or equal to 5")
	}

	// Exercise 5: Use a for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Exercise 6: Use a switch statement
	grade := 85
	switch {
	case grade >= 90:
		fmt.Println("Grade A")
	case grade >= 80:
		fmt.Println("Grade B")
	case grade >= 70:
		fmt.Println("Grade C")
	case grade >= 60:
		fmt.Println("Grade D")
	default:
		fmt.Println("Grade F")
	}
}