package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：交换两个变量的值
	a := 5
	b := 10
	a, b = b, a
	fmt.Println("a =", a, "b =", b)

	// 练习3：求和函数
	num1 := 5
	num2 := 10
	sum := add(num1, num2)
	fmt.Println("Sum:", sum)
}

func add(a, b int) int {
	return a + b
}