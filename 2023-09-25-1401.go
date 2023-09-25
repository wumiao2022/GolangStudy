package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：定义一个变量并打印
	var num int = 10
	fmt.Println(num)

	// 练习3：for循环打印数字
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 练习4：切片操作
	numbers := []int{1, 2, 3, 4, 5}
	slice := numbers[1:3]
	fmt.Println(slice)

	// 练习5：函数调用
	result := add(3, 5)
	fmt.Println(result)
}

// 练习5：函数定义
func add(a, b int) int {
	return a + b
}
