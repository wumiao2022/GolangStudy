package main

import (
	"fmt"
)

func main() {
	// 1. 打印出 "Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 学习变量的声明与赋值
	var name string = "John"
	age := 25
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// 3. 学习数组的使用
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	fmt.Println(numbers)

	// 4. 学习条件判断语句
	score := 85
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}

	// 5. 学习循环语句
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 6. 学习函数的定义与调用
	result := add(2, 3)
	fmt.Println(result)
}

// 定义一个函数，计算两个数的和
func add(a, b int) int {
	return a + b
}
