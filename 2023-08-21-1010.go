package main

import "fmt"

func main() {
	// 练习1：输出Hello World
	fmt.Println("Hello World!")

	// 练习2：计算并输出1+2的结果
	result := 1 + 2
	fmt.Println(result)

	// 练习3：定义一个变量name，并将其赋值为"Go"
	name := "Go"
	fmt.Println(name)

	// 练习4：定义一个整型数组，包含1, 2, 3三个元素
	numbers := [...]int{1, 2, 3}
	fmt.Println(numbers)

	// 练习5：定义一个函数add，接受两个整型参数x和y，并返回它们的和
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(2, 3))
}