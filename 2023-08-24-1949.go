package main

import (
	"fmt"
	"time"
)

func main() {
	// 当前日期
	today := time.Now().Format("2006-01-02")

	// 打印日期
	fmt.Println("Today is", today)

	// 打印Hello, World!
	fmt.Println("Hello, World!")

	// 定义一个整数变量
	num := 42

	// 打印变量的值
	fmt.Println("The value of num is:", num)

	// 定义并打印一个字符串数组
	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Println("Names:", names)

	// 使用for循环打印数组元素
	fmt.Println("Array elements:")
	for i, name := range names {
		fmt.Println(i, "-", name)
	}

	// 定义一个函数，计算两个整数的和，并返回结果
	add := func(a, b int) int {
		return a + b
	}

	// 调用函数并打印结果
	fmt.Println("The sum of 2 and 3 is:", add(2, 3))

	// 定义一个结构体类型
	type Person struct {
		Name string
		Age  int
	}

	// 创建结构体实例并打印
	person := Person{Name: "Alice", Age: 20}
	fmt.Println("Person:", person)
}