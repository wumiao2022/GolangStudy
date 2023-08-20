package main

import (
	"fmt"
)

func main() {
	// 1. 将以下字符串转换为字节切片，并打印出来
	str := "Golang"
	bytes := []byte(str)
	fmt.Println(bytes)

	// 2. 使用for循环遍历数组[1,2,3,4,5]，并打印出每个元素
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 3. 定义一个结构体Person，包含name和age两个字段，并创建一个Person结构体实例并打印出来
	type Person struct {
		name string
		age  int
	}

	person := Person{
		name: "John",
		age:  25,
	}

	fmt.Println(person)
}