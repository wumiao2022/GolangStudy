package main

import "fmt"

func main() {
	// 练习1: 打印Hello World
	fmt.Println("Hello World")

	// 练习2: 输出1到10的奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}

	// 练习3: 计算两个数的和
	num1, num2 := 5, 7
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习4: 判断一个数是否为正数
	num := -3
	if num > 0 {
		fmt.Println(num, "is positive")
	} else {
		fmt.Println(num, "is not positive")
	}

	// 练习5: 输出斐波那契数列前10个数
	n := 10
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}
}
