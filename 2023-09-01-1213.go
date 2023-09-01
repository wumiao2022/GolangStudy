package main

import "fmt"

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 计算两个整数的和
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3: 将1到10的整数相加
	total := 0
	for i := 1; i <= 10; i++ {
		total += i
	}
	fmt.Println("Total:", total)

	// 练习4: 判断一个数是否是偶数
	num := 9
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}