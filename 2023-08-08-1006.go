package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算并打印1到10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习3：交换两个变量的值
	a := 10
	b := 20
	fmt.Println("Before swap - a:", a, "b:", b)
	a, b = b, a
	fmt.Println("After swap - a:", a, "b:", b)

	// 练习4：使用递归计算斐波那契数列
	fmt.Println("Fibonacci sequence:")
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci(i))
	}

	// 练习5：找出切片中的最大值
	numbers := []int{12, 45, 2, 67, 89, 54}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	fmt.Println("Max number:", max)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
