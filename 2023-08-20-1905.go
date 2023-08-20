package main

import "fmt"

func main() {
	// Example 1
	fmt.Println("Hello, World!")

	// Example 2
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", sum)

	// Example 3
	var x string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&x)
	fmt.Println("You entered:", x)

	// Example 4
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println("Array:", arr)

	// Example 5
	slice := make([]int, 3, 5)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	fmt.Println("Slice:", slice)

	// Example 6
	m := make(map[string]int)
	m["apple"] = 1
	m["banana"] = 2
	m["cherry"] = 3
	fmt.Println("Map:", m)

	// Example 7
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// Example 8
	i := 0
	for i < 5 {
		i++
		fmt.Println("Iteration:", i)
	}

	// Example 9
	names := []string{"Alice", "Bob", "Charlie"}
	for index, value := range names {
		fmt.Println("Index:", index, "Value:", value)
	}

	// Example 10
	func add(a, b int) int {
		return a + b
	}
	fmt.Println("Sum:", add(3, 5))
}