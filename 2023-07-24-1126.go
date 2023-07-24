package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：两数相加
	num1 := 5
	num2 := 3
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断奇偶数
	num := 7
	if num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

	// 练习4：使用for循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：定义并使用函数
	num3 := 6
	result := square(num3)
	fmt.Println("Square:", result)
}

// 函数：计算一个数的平方
func square(n int) int {
	return n * n
}