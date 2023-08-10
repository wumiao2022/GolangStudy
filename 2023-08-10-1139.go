package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：打印1到10之间的所有奇数
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 练习3：判断一个数是否为质数
	num := 17 // 需要判断的数
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println("是质数")
	} else {
		fmt.Println("不是质数")
	}

	// 练习4：实现一个简单的计算器，可以进行加减乘除运算
	operator := "+" // 操作符，支持"+", "-", "*", "/"
	num1 := 10     // 第一个操作数
	num2 := 5      // 第二个操作数
	result := 0    // 运算结果
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		fmt.Println("无效的操作符")
		return
	}
	fmt.Println("计算结果为:", result)
}