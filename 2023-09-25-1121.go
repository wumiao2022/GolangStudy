package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：计算两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数
	num := 25
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：使用递归计算阶乘
	num := 5
	factorial := calculateFactorial(num)
	fmt.Println("Factorial of", num, "is", factorial)
}

func calculateFactorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * calculateFactorial(num-1)
}
