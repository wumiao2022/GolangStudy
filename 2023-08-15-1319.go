package main

import (
	"fmt"
)

func main() {
	// 1. 打印 "Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 声明一个整数变量，并打印出其值
	var num int = 10
	fmt.Println(num)

	// 3. 声明两个字符串变量，并将它们拼接在一起
	str1 := "Hello,"
	str2 := " World!"
	fmt.Println(str1 + str2)

	// 4. 声明一个切片，并向其中添加四个元素
	slice := []int{1, 2, 3, 4}
	fmt.Println(slice)

	// 5. 声明一个映射，并向其中添加三个键值对
	mapping := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	fmt.Println(mapping)
}