package main

import "fmt"

func main() {
	// 打印Hello, World!
	fmt.Println("Hello, World!")

	// 定义整型变量a和b，计算它们的和并输出结果
	a := 1
	b := 2
	c := a + b
	fmt.Println(c)

	// 使用for循环输出0~9的数字
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 定义一个切片，往其中添加元素并打印结果
	s := []int{1, 2, 3}
	s = append(s, 4)
	s = append(s, 5, 6)
	fmt.Println(s)

	// 定义一个函数，计算两个整数的和
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(1, 2))
}