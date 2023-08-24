package main

import (
	"fmt"
)

func main() {
	// 实例1：打印"Hello, World!"
	fmt.Println("Hello, World!")

	// 实例2：计算两个整数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 实例3：判断一个数是否为偶数
	num := 16
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 实例4：遍历切片
	nums := []int{1, 2, 3, 4, 5}
	for _, num := range nums {
		fmt.Println(num)
	}

	// 实例5：使用函数实现阶乘
	fmt.Println(factorial(5))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}