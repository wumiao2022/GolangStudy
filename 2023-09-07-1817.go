package main

import "fmt"

func main() {
	// 实现一个计算器程序，接收两个数和一个操作符，根据操作符计算结果并输出
	var num1, num2 int
	var operator string

	fmt.Println("请输入两个数和一个操作符，用空格分隔：")
	fmt.Scanln(&num1, &num2, &operator)

	switch operator {
	case "+":
		fmt.Println(num1, "+", num2, "=", num1+num2)
	case "-":
		fmt.Println(num1, "-", num2, "=", num1-num2)
	case "*":
		fmt.Println(num1, "*", num2, "=", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Println(num1, "/", num2, "=", num1/num2)
		} else {
			fmt.Println("除数不能为0")
		}
	default:
		fmt.Println("无效的操作符")
	}
}