package main

import (
	"fmt"
)

func main() {
	// 练习示例1：打印Hello World
	fmt.Println("Hello World")

	// 练习示例2：定义一个整型变量并赋初值，打印该变量的值
	num := 10
	fmt.Println(num)

	// 练习示例3：使用for循环打印1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习示例4：定义一个包含整型元素的切片，并使用range遍历打印每个元素的值
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 练习示例5：定义一个以匿名函数作为参数的函数，并调用该函数
	funcParam(func() {
		fmt.Println("Hello from anonymous function")
	})
}

// 函数接受一个函数类型的参数，并调用该函数
func funcParam(f func()) {
	f()
}