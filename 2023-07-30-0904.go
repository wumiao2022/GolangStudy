package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：将两个整数相加并打印结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：使用循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4：定义一个切片，并输出切片的长度和容量
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))

	// 练习5：使用函数实现阶乘，并打印结果
	factorial := func(n int) int {
		if n <= 1 {
			return 1
		}
		return n * factorial(n-1)
	}
	fmt.Println("Factorial of 5:", factorial(5))
}