package main

import "fmt"

func main() {
	// 练习1: 打印"Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2: 定义一个整型变量x，并赋值为10
	x := 10

	// 练习3: 判断x是否等于10，如果是则打印"x等于10"
	if x == 10 {
		fmt.Println("x等于10")
	}

	// 练习4: 定义一个字符串切片，并初始化包含三个元素："apple", "banana", "cherry"
	fruits := []string{"apple", "banana", "cherry"}

	// 练习5: 遍历打印切片中的元素
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// 练习6: 定义一个函数add，接收两个整型参数，并返回它们的和
	add := func(a, b int) int {
		return a + b
	}

	// 练习7: 调用add函数，传入参数2和3，并将结果赋值给变量sum
	sum := add(2, 3)

	// 练习8: 打印变量sum的值
	fmt.Println(sum)
}