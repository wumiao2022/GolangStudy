package main

import (
	"fmt"
)

func main() {
	// 练习1：输出Hello World
	fmt.Println("Hello, World!")

	// 练习2：计算1+2的结果并输出
	fmt.Println(1 + 2)

	// 练习3：求出5的阶乘并输出
	fmt.Println(5*4*3*2*1)

	// 练习4：定义一个变量x为10，输出x的值
	x := 10
	fmt.Println(x)

	// 练习5：定义一个函数sum，输入两个参数并返回它们的和
	fmt.Println(sum(3, 5))
}

func sum(a, b int) int {
	return a + b
}