package main

import "fmt"

func main() {
	// 1. 输出"Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 声明一个整数变量并赋值为10，输出变量的值
	var num int = 10
	fmt.Println(num)

	// 3. 声明一个浮点数变量并赋值为3.14，输出变量的值
	var pi float64 = 3.14
	fmt.Println(pi)

	// 4. 声明一个布尔变量并赋值为true，输出变量的值
	var isTrue bool = true
	fmt.Println(isTrue)

	// 5. 声明一个字符串变量并赋值为"Go语言"，输出变量的值
	var lang string = "Go语言"
	fmt.Println(lang)

	// 6. 声明一个整数切片，并往切片中添加元素1、2、3，输出切片的值
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	// 7. 声明一个字典，并添加键值对"name":"Tom"，输出字典的值
	dict := map[string]string{
		"name": "Tom",
	}
	fmt.Println(dict)

	// 8. 声明一个函数add，接收两个整数参数，返回它们的和
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(3, 4))

	// 9. 使用for循环输出1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 10. 使用if-else判断，如果num大于等于10，输出"num大于等于10"，否则输出"num小于10"
	if num >= 10 {
		fmt.Println("num大于等于10")
	} else {
		fmt.Println("num小于10")
	}
}