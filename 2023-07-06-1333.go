package main

import (
	"fmt"
)

func main() {
	// 练习1：打印"Hello, World!"
	fmt.Println("Hello, World!")
	
	// 练习2：定义一个整型变量并赋值为10，打印该变量的值
	num := 10
	fmt.Println(num)
	
	// 练习3：定义一个字符串变量并赋值为"Hello, Golang!"，打印该变量的值
	str := "Hello, Golang!"
	fmt.Println(str)
	
	// 练习4：定义一个切片变量并初始化为[1, 2, 3, 4, 5]，打印该切片的长度和元素值
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(len(slice))
	fmt.Println(slice)
	
	// 练习5：定义一个函数add，接受两个整型参数并返回它们的和，调用函数并打印结果
	result := add(3, 4)
	fmt.Println(result)
}

func add(a, b int) int {
	return a + b
}