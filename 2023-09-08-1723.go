package main

import "fmt"

func main() {
	// 示例1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 示例2：计算两个整数之和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 示例3：判断一个数是奇数还是偶数
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 示例4：遍历数组并打印元素
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 示例5：使用函数计算阶乘
	fmt.Println(factorial(5))
}

// 定义一个阶乘函数
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
