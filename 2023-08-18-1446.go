package main

import "fmt"

func main() {
	// 实例1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 实例2：计算两个整数的和并输出
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 实例3：使用for循环输出1到10的奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}

	// 实例4：定义一个切片并遍历输出其中的元素
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 实例5：使用递归计算斐波那契数列的第n个数
	n := 6
	fibonacci := calcFibonacci(n)
	fmt.Println("Fibonacci:", fibonacci)
}

func calcFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return calcFibonacci(n-1) + calcFibonacci(n-2)
}