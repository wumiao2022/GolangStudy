package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：使用for循环打印1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：定义一个函数，将两个整数相加并返回结果
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(3, 4))

	// 练习4：定义一个结构体Person，包含姓名和年龄字段，并初始化一个Person对象
	type Person struct {
		Name string
		Age  int
	}
	p := Person{
		Name: "Alice",
		Age:  25,
	}
	fmt.Println(p)

	// 练习5：使用切片存储一组整数，并打印切片的长度和容量
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))
}