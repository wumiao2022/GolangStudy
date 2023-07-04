package main

import "fmt"

func main() {
	// Example 1: Hello World
	fmt.Println("Hello, World!")

	// Example 2: Variable declaration and initialization
	var num int
	num = 10
	fmt.Println(num)

	// Example 3: Arithmetic operations
	a := 5
	b := 3
	c := a + b
	fmt.Println(c)

	// Example 4: If-else statement
	if num > 5 {
		fmt.Println("Number is greater than 5")
	} else {
		fmt.Println("Number is less than or equal to 5")
	}

	// Example 5: Looping - for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Example 6: Array
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// Example 7: Slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// Example 8: Map
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)

	// Example 9: Function
	multiply(2, 3)

	// Example 10: Package import
	fmt.Println("Square root of 16 is:", mymath.Sqrt(16))
}

func multiply(a, b int) {
	result := a * b
	fmt.Println(result)
}
