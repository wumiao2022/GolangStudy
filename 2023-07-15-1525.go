package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数
	number := 15
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}

	// 练习4：使用循环打印出0到9的数字
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 练习5：定义一个函数，计算一个数的平方
	fmt.Println(square(5))
}

func square(num int) int {
	return num * num
}
