package main

import "fmt"

func main() {
	// 1. 声明一个整型变量并初始化为10
	num := 10
	
	// 2. 声明一个字符串变量并初始化为"Hello, World!"
	message := "Hello, World!"
	
	// 3. 声明一个浮点型变量并初始化为3.14
	pi := 3.14
	
	// 4. 声明一个布尔型变量并初始化为true
	isTrue := true
	
	// 5. 声明一个切片并初始化为包含1、2、3三个元素的切片
	slice := []int{1, 2, 3}
	
	// 6. 使用循环打印出切片中的每个元素
	for _, val := range slice {
		fmt.Println(val)
	}
	
	// 7. 使用条件语句判断num的值是否大于5，并打印出结果
	if num > 5 {
		fmt.Println("num is greater than 5")
	} else {
		fmt.Println("num is not greater than 5")
	}
	
	// 8. 声明一个函数并调用它
	sayHello()
}

// 定义一个函数，用于打印"Hello, Golang!"
func sayHello() {
	fmt.Println("Hello, Golang!")
}
