package main

import "fmt"

func main() {
	// Example 1: Print Hello, World!
	fmt.Println("Hello, World!")

	// Example 2: Addition of two numbers
	fmt.Println(2 + 3)

	// Example 3: Variable declaration and assignment
	var x int = 5
	var y int = 3
	fmt.Println(x + y)

	// Example 4: String concatenation
	var name string = "Alice"
	var age int = 25
	fmt.Println("My name is " + name + " and I am " + string(age) + " years old.")

	// Example 5: Looping through an array
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// Example 6: If-Else statement
	isRaining := true
	if isRaining {
		fmt.Println("Take an umbrella.")
	} else {
		fmt.Println("Enjoy the sunshine.")
	}

	// Example 7: Function with return value
	result := addNumbers(2, 3)
	fmt.Println(result)

	// Example 8: Struct and methods
	type Rectangle struct {
		width  int
		height int
	}
	rect := Rectangle{10, 5}
	fmt.Println(rect.area())

	// Example 9: Pointer usage
	num := 10
	increase(&num)
	fmt.Println(num)
}

func addNumbers(a, b int) int {
	return a + b
}

func (rect Rectangle) area() int {
	return rect.width * rect.height
}

func increase(num *int) {
	*num = *num + 1
}
