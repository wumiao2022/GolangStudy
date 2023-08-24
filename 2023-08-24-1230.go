package main

import "fmt"

func main() {
	// 练习1：输出Hello World!
	fmt.Println("Hello World!")

	// 练习2：定义一个变量x为整数类型，并给它赋值为10，打印出x的值
	var x int = 10
	fmt.Println(x)

	// 练习3：定义一个字符串变量name，给它赋值为你的名字，打印出"Hello, 名字！"
	name := "John"
	fmt.Println("Hello,", name, "!")

	// 练习4：定义一个切片numbers，包含一些整数值，统计切片的长度和容量并打印出来
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))

	// 练习5：定义一个函数add，接收两个整数参数，返回它们的和
	fmt.Println(add(3, 4))
}

func add(a, b int) int {
	return a + b
}