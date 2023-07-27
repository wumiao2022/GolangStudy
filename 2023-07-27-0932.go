package main

import "fmt"

func main() {
	// 1. 打印 "Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 声明一个整数变量，并将其初始化为10，打印该变量的值
	var num int = 10
	fmt.Println(num)

	// 3. 声明一个浮点数变量，并将其初始化为3.14，打印该变量的值
	var pi float64 = 3.14
	fmt.Println(pi)

	// 4. 声明一个字符串变量，并将其初始化为"Go语言"，打印该变量的值
	var lang string = "Go语言"
	fmt.Println(lang)

	// 5. 声明一个布尔变量，并将其初始化为true，打印该变量的值
	var flag bool = true
	fmt.Println(flag)

	// 6. 声明一个整数切片，并往切片中追加一些整数，打印切片内容和长度
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(len(nums))

	// 7. 使用循环语句计算1到10的和，并打印结果
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 8. 声明一个Map，键为字符串类型，值为整数类型，添加一些键值对，并打印Map内容和长度
	m := make(map[string]int)
	m["apple"] = 1
	m["banana"] = 2
	m["orange"] = 3
	fmt.Println(m)
	fmt.Println(len(m))
}