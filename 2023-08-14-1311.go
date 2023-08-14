package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：创建一个整型变量并打印其值
	num := 10
	fmt.Println(num)

	// 练习3：创建一个只接受两个整型参数并返回它们和的函数
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(5, 3))

	// 练习4：创建一个切片，并使用循环遍历打印切片中的元素
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 练习5：使用函数闭包创建一个计数器，并测试其功能
	counter := func() func() int {
		n := 0
		return func() int {
			n = n + 1
			return n
		}
	}()
	fmt.Println(counter()) // 输出1
	fmt.Println(counter()) // 输出2
	fmt.Println(counter()) // 输出3
}