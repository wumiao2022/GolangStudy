package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：定义一个整数变量并初始化为10，打印该变量的值
	num := 10
	fmt.Println(num)

	// 练习3：定义一个字符串变量并初始化为"Go is awesome!"，打印该变量的值
	str := "Go is awesome!"
	fmt.Println(str)

	// 练习4：定义一个切片变量并初始化为包含5个整数的切片，打印该切片的长度
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(len(slice))

	// 练习5：定义一个函数，接受两个整数参数并返回它们的和
	sum := add(3, 4)
	fmt.Println(sum)
}

func add(a, b int) int {
	return a + b
}