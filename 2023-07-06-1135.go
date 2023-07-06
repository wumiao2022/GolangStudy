package main

import (
	"fmt"
)

func main() {
	// 输出Hello, World!
	fmt.Println("Hello, World!")

	// 数组操作
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	// 切片操作
	names := []string{"Alice", "Bob", "Charlie", "Dave"}
	fmt.Println(names)

	// 循环操作
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 条件判断
	num := 10
	if num > 5 {
		fmt.Println("Number is greater than 5")
	}

	// 函数调用
	result := add(2, 3)
	fmt.Println(result)
}

// 定义一个加法函数
func add(a, b int) int {
	return a + b
}