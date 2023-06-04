package main

import (
	"fmt"
)

func main() {
	//练习一：使用for循环打印1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//练习二：使用if语句判断2个数的大小关系
	num1 := 10
	num2 := 5
	if num1 > num2 {
		fmt.Printf("%d大于%d\n", num1, num2)
	} else if num1 < num2 {
		fmt.Printf("%d小于%d\n", num1, num2)
	} else {
		fmt.Println("两数相等")
	}

	//练习三：使用switch语句完成简单的计算器功能
	operator := "+"
	num3 := 10
	num4 := 5
	switch operator {
	case "+":
		fmt.Printf("%d+%d=%d", num3, num4, num3+num4)
	case "-":
		fmt.Printf("%d-%d=%d", num3, num4, num3-num4)
	case "*":
		fmt.Printf("%d*%d=%d", num3, num4, num3*num4)
	case "/":
		fmt.Printf("%d/%d=%d", num3, num4, num3/num4)
	default:
		fmt.Println("未知的操作符")
	}
	
	//练习四：使用for循环打印出九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}