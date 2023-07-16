package main

import (
	"fmt"
)

func main() {
	// 1. 打印"Hello World!"
	fmt.Println("Hello World!")

	// 2. 声明一个整数变量并赋值为10，打印该变量的值
	var num1 int = 10
	fmt.Println(num1)

	// 3. 声明一个字符串变量并赋值为"GoLang"，打印该变量的值
	var str1 string = "GoLang"
	fmt.Println(str1)

	// 4. 声明一个浮点数变量并赋值为3.14，打印该变量的值
	var float1 float64 = 3.14
	fmt.Println(float1)

	// 5. 声明一个布尔类型变量并赋值为true，打印该变量的值
	var bool1 bool = true
	fmt.Println(bool1)

	// 6. 声明一个数组，并初始化为[1, 2, 3, 4, 5]，打印该数组的值
	arr1 := [5]int {1, 2, 3, 4, 5}
	fmt.Println(arr1)

	// 7. 声明一个切片，并将数组[1, 2, 3, 4, 5]的前三个元素赋值给切片，打印该切片的值
	slice1 := arr1[:3]
	fmt.Println(slice1)

	// 8. 声明一个map，并初始化为{"name": "Bob", "age": 18}，打印该map的值
	map1 := map[string]interface{}{"name": "Bob", "age": 18}
	fmt.Println(map1)

	// 9. 声明一个函数，实现两个整数相加的功能，并打印相加结果
	add := func(a, b int) {
		fmt.Println(a + b)
	}
	add(2, 3)

	// 10. 声明一个结构体，表示一个人的信息，包括姓名和年龄
	type Person struct {
		Name string
		Age  int
	}
	person1 := Person{Name: "Alice", Age: 20}
	fmt.Println(person1)
}