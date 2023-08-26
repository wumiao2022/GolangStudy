package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：使用循环打印1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：使用if判断一个数是否为偶数，并打印结果
	n := 6
	if n%2 == 0 {
		fmt.Println(n, "是偶数")
	} else {
		fmt.Println(n, "不是偶数")
	}

	// 练习4：使用switch判断一个运算符，并打印对应的操作
	operator := "+"
	switch operator {
	case "+":
		fmt.Println("加法操作")
	case "-":
		fmt.Println("减法操作")
	case "*":
		fmt.Println("乘法操作")
	case "/":
		fmt.Println("除法操作")
	default:
		fmt.Println("未知操作")
	}

	// 练习5：定义一个数组并打印其中所有元素的和
	arr := [5]int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range arr {
		sum += num
	}
	fmt.Println("数组元素之和为:", sum)
}