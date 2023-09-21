package main

import (
	"fmt"
)

func main() {
	// Exercise 1
	fmt.Println("Hello, World!")

	// Exercise 2
	a := 10
	b := 5
	fmt.Println(a + b)

	// Exercise 3
	x := 5
	y := 7
	fmt.Println(x * y)

	// Exercise 4
	s := "Golang"
	fmt.Println("Welcome to", s)

	// Exercise 5
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("Sum:", sum)

	// Exercise 6
	numbers2 := []int{10, 20, 30, 40, 50}
	for i, num := range numbers2 {
		fmt.Println("Index:", i, "Number:", num)
	}

	// Exercise 7
	numbers3 := []int{2, 4, 6, 8, 10}
	total := 1
	for _, num := range numbers3 {
		total *= num
	}
	fmt.Println("Total:", total)
}