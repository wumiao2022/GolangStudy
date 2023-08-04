package main

import "fmt"

func main() {
	// 练习1: 打印 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 声明一个整型变量并打印出其值
	var num int
	fmt.Println(num)

	// 练习3: 声明一个字符串变量并打印出其值
	var str string
	fmt.Println(str)

	// 练习4: 声明一个切片变量并打印出其长度和容量
	var slice []int
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	// 练习5: 声明一个函数并调用它
	greet()

	// 练习6: 循环打印从1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// 练习5: 打印问候语
func greet() {
	fmt.Println("Hello!")
}