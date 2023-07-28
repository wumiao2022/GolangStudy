package main

import "fmt"

func main() {
	// 练习一：打印Hello World
	fmt.Println("Hello World")

	// 练习二：计算两个整数之和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习三：判断一个数是否为偶数
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习四：定义一个数组，并遍历输出数组中的元素
	arr := [5]int{1, 2, 3, 4, 5}
	for _, value := range arr {
		fmt.Println(value)
	}

	// 练习五：计算斐波那契数列的第n项
	n := 10
	fmt.Println("Fibonacci series:")
	for i := 0; i < n; i++ {
		fmt.Print(fibonacci(i), " ")
	}
	fmt.Println()
}

// 计算斐波那契数列的第n项
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}