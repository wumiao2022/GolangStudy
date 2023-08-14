package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算1+2的和并打印结果
	sum := 1 + 2
	fmt.Println("1 + 2 =", sum)

	// 练习3：定义一个slice并打印出其中的元素
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", numbers)

	// 练习4：使用for循环打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：定义一个函数，输入两个整数并返回它们的和
	add := func(a, b int) int {
		return a + b
	}

	result := add(3, 5)
	fmt.Println("3 + 5 =", result)
}