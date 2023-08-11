package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：计算并打印两个整数相加的结果
	a := 10
	b := 20
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习3：判断一个整数是否为偶数
	num := 45
	isEven := num%2 == 0
	fmt.Println("Is even:", isEven)

	// 练习4：使用循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：使用条件判断打印数字是否为正数、零还是负数
	number := -15
	if number > 0 {
		fmt.Println("Positive")
	} else if number == 0 {
		fmt.Println("Zero")
	} else {
		fmt.Println("Negative")
	}
}