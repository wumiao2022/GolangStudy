package main

import (
	"fmt"
)

func main() {
	// 练习1：输出"Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：计算并输出1+2的结果
	fmt.Println(1 + 2)

	// 练习3：定义一个变量x，并将其赋值为10，然后输出x的值
	x := 10
	fmt.Println(x)

	// 练习4：定义一个切片，包含元素1, 2, 3，并输出切片的长度
	slice := []int{1, 2, 3}
	fmt.Println(len(slice))

	// 练习5：定义一个函数add，接受两个整数参数，并返回它们的和
	fmt.Println(add(2, 3))
}

func add(a, b int) int {
	return a + b
}