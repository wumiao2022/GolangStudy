package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：定义一个整型变量并打印出其值
	num := 100
	fmt.Println(num)

	// 练习3：定义一个字符串变量并打印出其值
	str := "Golang"
	fmt.Println(str)

	// 练习4：定义一个布尔型变量并打印出其值
	flag := true
	fmt.Println(flag)

	// 练习5：定义一个切片并打印出其长度和容量
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// 练习6：使用for循环打印1到10之间的所有偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习7：定义一个函数，接收两个整型参数并返回它们的和
	sum := add(3, 5)
	fmt.Println(sum)
}

func add(a, b int) int {
	return a + b
}