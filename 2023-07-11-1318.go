package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：求和函数
	sum := add(2, 3)
	fmt.Println("Sum:", sum)

	// 练习3：判断奇偶数
	num := 5
	if isEven(num) {
		fmt.Println(num, "is an even number.")
	} else {
		fmt.Println(num, "is an odd number.")
	}

	// 练习4：切片操作
	slice := []string{"apple", "banana", "orange", "grape"}
	fmt.Println(slice[1:3])
}

func add(a, b int) int {
	return a + b
}

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}