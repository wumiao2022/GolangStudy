package main

import (
	"fmt"
)

func main() {
	// 1. Print "Hello, World!" to the console.
	fmt.Println("Hello, World!")

	// 2. Create a variable "name" and assign it a string value. Print the value of "name" to the console.
	name := "John Doe"
	fmt.Println(name)

	// 3. Create two variables "a" and "b" and assign them integer values. Print the sum of "a" and "b" to the console.
	a := 10
	b := 5
	sum := a + b
	fmt.Println(sum)

	// 4. Create an array with 5 string elements. Print each element to the console.
	array := [5]string{"apple", "banana", "cherry", "date", "elderberry"}
	for _, element := range array {
		fmt.Println(element)
	}

	// 5. Create a slice with 3 string elements. Append another string element to the slice. Print the slice to the console.
	slice := []string{"elephant", "flamingo", "giraffe"}
	slice = append(slice, "hippo")
	fmt.Println(slice)
}