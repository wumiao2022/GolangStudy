package main

import "fmt"

func main() {
	// 练习1: 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		fmt.Println()
	}

	// 练习2: 判断一个数字是否为素数
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Printf("%d is a prime number\n", num)
	} else {
		fmt.Printf("%d is not a prime number\n", num)
	}

	// 练习3: 实现一个简单的计算器函数，接收两个数字和一个运算符，并返回运算结果
	calculate(10, 5, "*")
	calculate(10, 5, "+")
	calculate(10, 5, "-")
	calculate(10, 5, "/")
}

func calculate(a, b int, operator string) {
	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		fmt.Println("Invalid operator")
		return
	}
	fmt.Printf("%d %s %d = %d\n", a, operator, b, result)
}