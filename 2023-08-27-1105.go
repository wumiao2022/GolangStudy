package main

import (
	"fmt"
)

func main() {
	// Example 1
	fmt.Println("Hello, World!")

	// Example 2
	x := 5
	y := 10
	sum := x + y
	fmt.Println(sum)

	// Example 3
	nums := []int{1, 2, 3, 4, 5}
	for _, num := range nums {
		fmt.Println(num)
	}

	// Example 4
	var a int = 20
	var ptr *int
	ptr = &a
	fmt.Println(*ptr)

	// Example 5
	var name string = "Golang"
	fmt.Println("Length:", len(name))
}