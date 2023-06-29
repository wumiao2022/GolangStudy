package main

import "fmt"

func main() {
	// 1. 声明一个整型变量x，并给它赋值为25
	x := 25

	// 2. 声明一个字符串变量name，并给它赋值为"John"
	name := "John"

	// 3. 创建一个整型切片nums，并初始化它为包含元素1、2、3的切片
	nums := []int{1, 2, 3}

	// 4. 使用for循环打印出nums切片中的每个元素
	for _, num := range nums {
		fmt.Println(num)
	}

	// 5. 声明一个匿名函数，该函数接受两个整型参数，返回它们的和
	sum := func(a, b int) int {
		return a + b
	}

	// 6. 调用匿名函数，并将结果打印出来
	fmt.Println(sum(3, 4))
}