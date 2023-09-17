package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, World!")

	// 打印当前时间
	fmt.Println(time.Now())

	// 数组声明和初始化
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// 循环打印数组元素
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 切片声明和初始化
	slice := []string{"apple", "banana", "cherry"}
	fmt.Println(slice)

	// 循环打印切片元素
	for _, item := range slice {
		fmt.Println(item)
	}

	// 声明和初始化一个映射
	m := make(map[string]int)
	m["apple"] = 2
	m["banana"] = 3
	m["cherry"] = 5
	fmt.Println(m)

	// 遍历映射并打印键值对
	for k, v := range m {
		fmt.Printf("%s: %d\n", k, v)
	}
}