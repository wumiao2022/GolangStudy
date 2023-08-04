package main

import "fmt"

func main() {
	// 1. 声明一个整数变量，初始化为10，将其输出到控制台
	var num int = 10
	fmt.Println(num)

	// 2. 声明一个字符串变量，初始化为"Hello, World!"，将其输出到控制台
	var str string = "Hello, World!"
	fmt.Println(str)

	// 3. 声明一个浮点数变量，初始化为3.14，将其输出到控制台
	var floatNum float64 = 3.14
	fmt.Println(floatNum)

	// 4. 声明一个布尔变量，初始化为true，将其输出到控制台
	var isTrue bool = true
	fmt.Println(isTrue)

	// 5. 声明一个整数数组，包含元素1, 2, 3，将其输出到控制台
	var nums = []int{1, 2, 3}
	fmt.Println(nums)

	// 6. 声明一个字符串数组，包含元素"apple", "banana", "orange"，将其输出到控制台
	var fruits = []string{"apple", "banana", "orange"}
	fmt.Println(fruits)

	// 7. 声明一个包含键值对的字典，键为整数类型，值为字符串类型，包含一对键值对{1: "one"}，将其输出到控制台
	var dict = map[int]string{1: "one"}
	fmt.Println(dict)
}