package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variables and Constants
	var num1 int = 10
	var num2 float64 = 3.14
	const num3 int = 5
	
	fmt.Println("Number 1:", num1)
	fmt.Println("Number 2:", num2)
	fmt.Println("Number 3:", num3)

	// Exercise 3: Arithmetic Operations
	result1 := num1 + int(num2)
	result2 := num1 - int(num2)
	result3 := num1 * int(num2)
	result4 := num1 / int(num2)
	result5 := num1 % int(num2)
	
	fmt.Println("Result 1:", result1)
	fmt.Println("Result 2:", result2)
	fmt.Println("Result 3:", result3)
	fmt.Println("Result 4:", result4)
	fmt.Println("Result 5:", result5)

	// Exercise 4: Conditional Statements
	if num1 > 0 {
		fmt.Println("Number 1 is positive")
	} else if num1 < 0 {
		fmt.Println("Number 1 is negative")
	} else {
		fmt.Println("Number 1 is zero")
	}

	if num2 > float64(num3) {
		fmt.Println("Number 2 is greater than Number 3")
	} else if num2 < float64(num3) {
		fmt.Println("Number 2 is smaller than Number 3")
	} else {
		fmt.Println("Number 2 is equal to Number 3")
	}

	// Exercise 5: Loops
	for i := 0; i < num1; i++ {
		fmt.Printf("Iteration %d\n", i)
	}

	j := 0
	for j < num1 {
		fmt.Printf("Iteration %d\n", j)
		j++
	}

	k := 0
	for {
		if k >= num1 {
			break
		}
		fmt.Printf("Iteration %d\n", k)
		k++
	}
}