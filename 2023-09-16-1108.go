package main

import "fmt"

func main() {
	// 1. 打印 "Hello, World!" 到控制台
	fmt.Println("Hello, World!")

	// 2. 定义并打印一个整数变量
	num := 10
	fmt.Println(num)

	// 3. 定义并打印一个字符串变量
	str := "Golang is awesome!"
	fmt.Println(str)

	// 4. 定义一个切片并打印其中的元素
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// 5. 使用循环打印出切片中的每个元素
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// 6. 定义一个函数，接收一个整数参数，并返回参数加上10的结果
	addTen := func(num int) int {
		return num + 10
	}

	// 7. 调用函数并打印结果
	result := addTen(5)
	fmt.Println(result)

	// 8. 定义一个结构体类型，并定义对应的变量
	type Person struct {
		Name string
		Age  int
	}

	person := Person{
		Name: "John",
		Age:  30,
	}

	// 9. 打印结构体变量的字段值
	fmt.Println(person.Name)
	fmt.Println(person.Age)
}