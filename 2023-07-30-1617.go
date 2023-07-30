package main

import "fmt"

func main() {
	// 练习1：输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个数的和并输出结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：循环输出1到10的数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：定义一个函数，计算两个数的差并返回结果
	fmt.Println("Difference:", subtract(20, 10))
}

func subtract(num1, num2 int) int {
	return num1 - num2
}