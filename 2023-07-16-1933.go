package main

import "fmt"

func main() {
	// 练习1：打印输出Hello World
	fmt.Println("Hello World")

	// 练习2：计算两个整数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数
	num := 42
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：使用for循环打印出1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：定义一个函数，计算两个浮点数的乘积
	multiply(2.5, 4.3)
}

func multiply(a, b float64) {
	product := a * b
	fmt.Println("Product:", product)
}