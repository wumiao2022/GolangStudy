package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：变量声明和赋值
	var num1 int = 10
	var num2 float64 = 3.14
	fmt.Println(num1, num2)

	// 练习3：判断语句
	if num1 > 5 {
		fmt.Println("num1 is greater than 5")
	} else {
		fmt.Println("num1 is not greater than 5")
	}

	// 练习4：循环语句
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration", i+1)
	}

	// 练习5：切片操作
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers[1:3])

	// 练习6：函数定义和调用
	result := add(3, 5)
	fmt.Println("Result:", result)
}

// 函数：两个整数相加
func add(a, b int) int {
	return a + b
}