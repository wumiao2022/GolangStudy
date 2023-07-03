package main

import "fmt"

func main() {
	// 练习1：打印 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：定义一个整型变量并进行加法计算
	var a, b int = 5, 10
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习3：使用 if-else 判断两个数的关系
	num1, num2 := 10, 15
	if num1 > num2 {
		fmt.Println("num1 is greater than num2")
	} else if num1 < num2 {
		fmt.Println("num1 is less than num2")
	} else {
		fmt.Println("num1 and num2 are equal")
	}

	// 练习4：使用 for 循环打印数字 1 到 5
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 练习5：使用 switch-case 判断一个数字的奇偶性
	number := 7
	switch number % 2 {
	case 0:
		fmt.Println("Even number")
	default:
		fmt.Println("Odd number")
	}
}