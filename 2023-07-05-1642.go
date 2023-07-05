package main

import "fmt"

func main() {
	// 练习1
	fmt.Println("Hello, World!")

	// 练习2
	var num1, num2 int = 10, 20
	fmt.Println(num1 + num2)

	// 练习3
	var x, y = "Hello", "World!"
	fmt.Println(x, y)

	// 练习4
	fmt.Println(add(5, 3))

	// 练习5
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}
}

func add(a, b int) int {
	return a + b
}
