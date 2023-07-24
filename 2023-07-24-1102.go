package main

import "fmt"

func main() {
	// 1. 声明一个整型变量并初始化为10，将其输出到控制台
	var num int = 10
	fmt.Println(num)

	// 2. 声明一个字符串变量并初始化为"Hello, Golang!"，将其输出到控制台
	var greeting string = "Hello, Golang!"
	fmt.Println(greeting)

	// 3. 定义一个包含5个元素的整型数组，并将其打印到控制台
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// 4. 声明一个切片，并添加5个整型元素，然后将其打印到控制台
	var slice []int
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Println(slice)

	// 5. 声明一个map，并添加3个键值对，然后将其打印到控制台
	var people map[string]int = map[string]int{
		"John":  25,
		"Sarah": 30,
		"Tom":   35,
	}
	fmt.Println(people)
}