package main

import "fmt"

func main() {
	// 实例1: 打印Hello World
	fmt.Println("Hello World")

	// 实例2: 定义和使用变量
	var num1 int = 10
	num2 := 20
	fmt.Println(num1 + num2)

	// 实例3: 使用循环
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 实例4: 使用条件语句
	age := 18
	if age >= 18 {
		fmt.Println("成年人")
	} else {
		fmt.Println("未成年人")
	}

	// 实例5: 使用切片
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// 实例6: 使用函数
	result := add(3, 5)
	fmt.Println(result)
}

func add(a, b int) int {
	return a + b
}