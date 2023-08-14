package main

import (
	"fmt"
)

func main() {
	// 练习1：打印 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：变量赋值与打印
	var num1 int = 10
	var num2 float64 = 3.14
	var str string = "Golang"
	fmt.Printf("num1: %d, num2: %.2f, str: %s\n", num1, num2, str)

	// 练习3：使用循环打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4：判断数字是偶数还是奇数
	num := 9
	if num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}

	// 练习5：使用函数计算两个数之和
	sum := add(5, 7)
	fmt.Println("Sum:", sum)
}

// 函数：计算两个数之和
func add(a, b int) int {
	return a + b
}