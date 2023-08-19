package main

import (
	"fmt"
)

func main() {
	// 1. 打印出"Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 声明一个变量x，赋值为10，并打印出x的值
	x := 10
	fmt.Println(x)

	// 3. 声明一个切片slice，包含整数1、2、3，并打印出slice的长度和容量
	slice := []int{1, 2, 3}
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	// 4. 声明一个包含键值对的map，包括姓名和年龄，并打印出map中姓名为"John"的年龄
	ages := map[string]int{
		"John":  25,
		"Peter": 30,
	}
	fmt.Println("John's age:", ages["John"])

	// 5. 使用循环语句输出1到10之间的所有奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}
}