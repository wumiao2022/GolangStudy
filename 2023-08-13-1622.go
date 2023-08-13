package main

import (
	"fmt"
)

func main() {
	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")

	// 2. 定义并打印变量
	var num1 int = 10
	var num2 float64 = 3.14
	fmt.Printf("num1: %d, num2: %.2f\n", num1, num2)

	// 3. 数组遍历
	arr := []int{1, 2, 3, 4, 5}
	for _, num := range arr {
		fmt.Println(num)
	}

	// 4. 条件判断
	if num1 > 0 {
		fmt.Println("num1是正数")
	} else {
		fmt.Println("num1是负数")
	}

	// 5. 循环语句
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// 6. 函数定义和调用
	sum := add(2, 3)
	fmt.Println("Sum:", sum)
}

// add函数，接受两个参数并返回它们的和
func add(a, b int) int {
	return a + b
}
