package main

import "fmt"

func main() {
	// Exercise 1
	fmt.Println("Hello, World!")

	// Exercise 2
	x := 10
	y := 5
	z := x + y
	fmt.Println(z)

	// Exercise 3
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 4
	numbers := []int{2, 4, 6, 8, 10}
	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}

	// Exercise 5
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum of numbers from 1 to 100:", sum)
}