package main

import "fmt"

func main() {
	// 1. 打印 Hello World
	fmt.Println("Hello World")

	// 2. 定义一个整数变量并打印
	var num int = 10
	fmt.Println(num)

	// 3. 定义一个字符串变量并打印
	var str string = "Golang"
	fmt.Println(str)

	// 4. 判断两个整数是否相等，并打印结果
	a := 5
	b := 7
	fmt.Println(a == b)

	// 5. 循环打印数字 1 到 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 6. 定义一个数组，并循环打印每个元素
	arr := [5]int{1, 2, 3, 4, 5}
	for _, v := range arr {
		fmt.Println(v)
	}

	// 7. 定义一个切片，并循环打印每个元素
	slice := []string{"a", "b", "c", "d", "e"}
	for _, v := range slice {
		fmt.Println(v)
	}

	// 8. 定义一个函数，接收两个整数参数，并返回它们的和
	sum := add(2, 3)
	fmt.Println(sum)
}

func add(a, b int) int {
	return a + b
}
