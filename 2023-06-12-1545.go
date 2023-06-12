package main

import (
	"fmt"
)

func main() {
	// 练习1：输出 "Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：输出 1~10 的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：求 1~10 的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习4：实现一个简单的计算器，输入两个数和运算符，输出计算结果
	var num1, num2 float64
	var operator string
	fmt.Print("请输入第一个数：")
	fmt.Scanln(&num1)
	fmt.Print("请输入第二个数：")
	fmt.Scanln(&num2)
	fmt.Print("请输入运算符：")
	fmt.Scanln(&operator)
	switch operator {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f", num1, num2, num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("除数不能为 0")
		} else {
			fmt.Printf("%.2f / %.2f = %.2f", num1, num2, num1/num2)
		}
	default:
		fmt.Println("不支持的操作符")
	}
}