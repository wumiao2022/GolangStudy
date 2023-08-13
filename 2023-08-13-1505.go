package main

import "fmt"

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 计算1+2的结果并打印
	sum := 1 + 2
	fmt.Println(sum)

	// 练习3: 遍历并打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4: 定义一个切片，并向其中添加元素后打印
	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4, 5)
	fmt.Println(numbers)

	// 练习5: 定义一个函数，接收两个整数参数并返回它们的和
	result := sumNumbers(3, 4)
	fmt.Println(result)
}

func sumNumbers(a, b int) int {
	return a + b
}