package main

import "fmt"

func main() {
	// 1. Hello World
	fmt.Println("Hello, World!")

	// 2. Variables
	var x int = 10
	fmt.Println(x)

	// 3. Constants
	const y = 5
	fmt.Println(y)

	// 4. If-else statement
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than or equal to 5")
	}

	// 5. For-loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 6. Arrays
	array := [3]int{1, 2, 3}
	fmt.Println(array)

	// 7. Slices
	slice := []int{4, 5, 6}
	fmt.Println(slice)

	// 8. Maps
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Println(m)

	// 9. Functions
	fmt.Println(add(3, 4))
}

func add(a, b int) int {
	return a + b
}