package main

import (
	"fmt"
)

func main() {
	// 练习1: 输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 计算两个整数的和并输出结果
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println(sum)

	// 练习3: 判断一个数是否为偶数并输出结果
	num := 7
	isEven := num%2 == 0
	fmt.Println(isEven)

	// 练习4: 使用循环求1到10的累加和并输出结果
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习5: 定义一个字符串变量并输出其长度
	str := "Hello, Golang!"
	fmt.Println(len(str))
}