package main

import "fmt"

func main() {
	// Example 1
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("Sum of numbers:", sum)

	// Example 2
	name := "Golang"
	length := len(name)
	reversedName := make([]rune, length)
	for i, char := range name {
		reversedName[length-1-i] = char
	}
	fmt.Println("Reversed name:", string(reversedName))

	// Example 3
	fruits := [3]string{"apple", "banana", "orange"}
	for _, fruit := range fruits {
		fmt.Println("I love", fruit)
	}
}