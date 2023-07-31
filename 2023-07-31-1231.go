package main

import "fmt"

func main() {
	// 练习1：声明一个整数变量x，并初始化为10，打印出x的值
	x := 10
	fmt.Println(x)

	// 练习2：声明两个整数变量a和b，并分别初始化为5和7，将它们相加并打印结果
	a := 5
	b := 7
	fmt.Println(a + b)

	// 练习3：声明一个字符串变量name并初始化为"GoLang"，打印出name的值和长度
	name := "GoLang"
	fmt.Println(name)
	fmt.Println(len(name))

	// 练习4：声明一个布尔变量isTrue并初始化为true，打印出isTrue的值
	isTrue := true
	fmt.Println(isTrue)

	// 练习5：声明一个切片变量numbers，并初始化为包含1、2、3三个元素的切片，打印出numbers的值和长度
	numbers := []int{1, 2, 3}
	fmt.Println(numbers)
	fmt.Println(len(numbers))
}