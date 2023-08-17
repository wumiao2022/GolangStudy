package main

import "fmt"

func main() {
	// 练习1
	fmt.Println("Hello, World!")

	// 练习2
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("The sum is:", sum)

	// 练习3
	age := 25
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// 练习4
	items := []string{"apple", "banana", "orange"}
	for _, item := range items {
		fmt.Println(item)
	}

	// 练习5
	fmt.Println(factorial(5))
}

// 练习5的辅助函数：计算阶乘
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}