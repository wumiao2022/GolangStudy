package main

import "fmt"

func main() {
	// 练习1: 打印 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 定义一个整型变量并赋值为10，打印出该变量的值
	num := 10
	fmt.Println(num)

	// 练习3: 定义一个切片，其中包含5个整数，并使用循环打印出切片中的值
	slice := []int{1, 2, 3, 4, 5}
	for _, num := range slice {
		fmt.Println(num)
	}

	// 练习4: 定义一个函数，接受两个整型参数并返回它们的和
	sum := add(3, 5)
	fmt.Println(sum)

	// 练习5: 定义一个结构体表示一个人的信息，包括姓名和年龄，并打印出该结构体的值
	person := Person{Name: "John", Age: 25}
	fmt.Println(person)
}

func add(a, b int) int {
	return a + b
}

type Person struct {
	Name string
	Age  int
}