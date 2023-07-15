package main

import "fmt"

func main() {
	// 打印 Hello, World!
	fmt.Println("Hello, World!")

	// 定义并打印一个整数变量
	num := 10
	fmt.Println(num)

	// 定义并打印一个字符串变量
	str := "Hello, Golang!"
	fmt.Println(str)

	// 定义一个切片并打印其中的元素
	slice := []int{1, 2, 3, 4, 5}
	for _, num := range slice {
		fmt.Println(num)
	}

	// 定义一个结构体，并打印其中的字段值
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "Alice", Age: 25}
	fmt.Println(person.Name, person.Age)
}