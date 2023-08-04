package main

import "fmt"

func main() {
	// 1. 使用 fmt 包打印输出 "Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 声明一个整数变量，赋值为 10，并打印输出
	var num1 int = 10
	fmt.Println(num1)

	// 3. 声明一个浮点数变量，赋值为 3.14，并打印输出
	var num2 float64 = 3.14
	fmt.Println(num2)

	// 4. 声明一个字符串变量，赋值为 "Golang"，并打印输出
	var str string = "Golang"
	fmt.Println(str)

	// 5. 使用循环输出数字 1 到 5
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 6. 判断一个数是否为正数，并打印输出结果
	num3 := -7
	if num3 > 0 {
		fmt.Println("正数")
	} else {
		fmt.Println("非正数")
	}

	// 7. 声明一个字符串切片，包含 "Golang", "Python", "Java"，并打印输出
	slice := []string{"Golang", "Python", "Java"}
	fmt.Println(slice)
}