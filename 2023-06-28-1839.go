package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 输出Hello, World!
	fmt.Println("Hello, World!")

	// 2. 输出当前时间
	fmt.Println("Current time:", time.Now().Format("2006-01-02 15:04:05"))

	// 3. 定义一个整型变量并输出
	num := 10
	fmt.Println("Number:", num)

	// 4. 求两个整数的和并输出
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 5. 判断一个数是否为偶数并输出结果
	number := 15
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}
}