package main

import "fmt"

func main() {
	// 示例：打印Hello, World!
	fmt.Println("Hello, World!")

	// 示例：计算并打印两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 示例：使用for循环打印1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 示例：判断一个数是否为偶数，并打印判断结果
	num := 24
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 示例：定义一个字符串切片，并遍历打印其中的元素
	strings := []string{"apple", "banana", "cherry"}
	for _, str := range strings {
		fmt.Println(str)
	}

	// 示例：使用函数计算两个数的乘积，并打印结果
	product := multiply(5, 6)
	fmt.Println("Product:", product)
}

// 函数：计算两个数的乘积
func multiply(a, b int) int {
	return a * b
}