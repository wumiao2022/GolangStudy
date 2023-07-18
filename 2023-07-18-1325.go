package main

import (
	"fmt"
)

func main() {
	// 1. 声明一个函数变量并调用
	f := func() {
		fmt.Println("Hello, Go!")
	}
	f()

	// 2. 使用循环计算1到10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 3. 定义一个结构体类型，并创建一个该类型的变量
	type Person struct {
		Name string
		Age  int
	}

	p := Person{
		Name: "Alice",
		Age:  25,
	}
	fmt.Println("Person:", p)

	// 4. 使用切片操作截取字符串的一部分
	str := "Hello, Go!"
	subStr := str[7:]
	fmt.Println("Substring:", subStr)

	// 5. 通过递归实现阶乘函数
	fmt.Println("Factorial(5):", factorial(5))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
