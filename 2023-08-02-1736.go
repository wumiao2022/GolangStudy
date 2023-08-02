package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World!")

	// 练习2：定义一个整型变量并赋值为10，打印该变量的值
	num := 10
	fmt.Println(num)

	// 练习3：定义一个字符串变量并赋值为"Hello"，并将其与另一个字符串变量"World"进行拼接，打印拼接后的结果
	str1 := "Hello"
	str2 := "World"
	fmt.Println(str1 + " " + str2)

	// 练习4：定义一个整型数组，并赋值为[1, 2, 3, 4, 5]，遍历数组并打印每个元素的值
	nums := [5]int{1, 2, 3, 4, 5}
	for _, value := range nums {
		fmt.Println(value)
	}

	// 练习5：定义一个函数，接收两个整型参数，将它们相加，并返回结果
	result := add(10, 20)
	fmt.Println(result)
}

func add(a, b int) int {
	return a + b
}