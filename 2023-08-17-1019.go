package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：定义一个整数变量x，并将其初始化为10，然后打印出x的值
	x := 10
	fmt.Println(x)

	// 练习3：声明一个字符串变量，并将其赋值为"Go语言是最好的编程语言！"，然后打印出字符串的长度
	str := "Go语言是最好的编程语言！"
	fmt.Println(len(str))

	// 练习4：定义一个切片，其中包含整数1到5的元素，并打印出切片的长度和容量
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(len(slice), cap(slice))

	// 练习5：定义一个函数，接收两个整数参数，返回它们的和
	sum := add(3, 5)
	fmt.Println(sum)
}

func add(a, b int) int {
	return a + b
}