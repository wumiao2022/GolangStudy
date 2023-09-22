package main

import (
	"fmt"
)

func main() {
	// 1. 输出 "Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 声明一个整数变量x并赋值为10，输出x的值
	x := 10
	fmt.Println(x)

	// 3. 声明一个字符串变量name并赋值为"John"，输出name的值
	name := "John"
	fmt.Println(name)

	// 4. 声明一个数组arr，包含元素1, 2, 3, 4, 5，并输出数组的长度和第三个元素的值
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(len(arr))
	fmt.Println(arr[2])

	// 5. 声明一个切片slice，包含元素"a", "b", "c"，并输出切片的长度和第一个元素的值
	slice := []string{"a", "b", "c"}
	fmt.Println(len(slice))
	fmt.Println(slice[0])
}