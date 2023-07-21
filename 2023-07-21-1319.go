package main

import "fmt"

func main() {
	// 实例1
	fmt.Println("Hello, World!")

	// 实例2
	var num1, num2 int
	num1 = 10
	num2 = 20
	fmt.Println("Sum:", num1+num2)

	// 实例3
	names := []string{"Alice", "Bob", "Charlie"}
	for _, name := range names {
		fmt.Println("Hello,", name)
	}

	// 实例4
	isTrue := true
	if isTrue {
		fmt.Println("It is true!")
	} else {
		fmt.Println("It is false!")
	}

	// 实例5
	var n int = 5
	if n > 0 {
		fmt.Println("Positive number!")
	} else if n < 0 {
		fmt.Println("Negative number!")
	} else {
		fmt.Println("Zero!")
	}
}