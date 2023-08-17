package main

import "fmt"

func main() {
	// Example 1: Using if-else statement
	number := 10
	if number%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	// Example 2: Using for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Example 3: Using switch statement
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	default:
		fmt.Println("Today is not Monday or Tuesday")
	}

	// Example 4: Using arrays
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// Example 5: Using slices
	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Println(names)
}