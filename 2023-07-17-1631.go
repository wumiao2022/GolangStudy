package main

import "fmt"

func main() {
	// 题目：输出Hello, World!
	fmt.Println("Hello, World!")

	// 题目：声明一个变量并输出其值
	var num int = 30
	fmt.Println(num)

	// 题目：声明一个字符串变量并输出其值
	var name string = "John"
	fmt.Println(name)

	// 题目：声明两个整数变量并输出其和
	var a int = 10
	var b int = 20
	sum := a + b
	fmt.Println(sum)

	// 题目：声明一个整数切片并输出其长度
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(len(numbers))

	// 题目：声明一个空接口，并将一个整数赋值给它
	var i interface{}
	i = 42
	fmt.Println(i)

	// 题目：声明一个Map，并向其中添加键值对，然后输出value值
	person := make(map[string]string)
	person["name"] = "Alice"
	person["age"] = "25"
	fmt.Println(person["age"])
}