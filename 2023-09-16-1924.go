package main

import (
	"fmt"
)

func main() {
	// 示例1：打印Hello World
	fmt.Println("Hello World!")

	// 示例2：定义一个整型变量并赋值，打印出变量的值
	var num int
	num = 10
	fmt.Println(num)

	// 示例3：定义一个字符串变量并赋值，打印出变量的值
	var str string
	str = "Golang"
	fmt.Println(str)

	// 示例4：使用for循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 示例5：创建一个切片，并遍历打印切片中的元素
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}
}