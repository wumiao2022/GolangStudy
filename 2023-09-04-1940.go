package main

import "fmt"

func main() {
	// 1. 打印10到1的数字，使用for循环
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}

	// 2. 声明一个切片，包含元素 "apple", "banana", "orange"
	fruits := []string{"apple", "banana", "orange"}
	fmt.Println(fruits)

	// 3. 声明一个map，将学生的姓名作为键，年龄作为值
	students := map[string]int{
		"Alice":  20,
		"Bob":    22,
		"Charlie": 19,
	}
	fmt.Println(students)

	// 4. 声明一个结构体类型Person，包含姓名和年龄字段
	type Person struct {
		Name string
		Age  int
	}
	
	// 5. 声明一个指针变量p，指向结构体Person的实例
	p := &Person{
		Name: "John",
		Age:  25,
	}
	fmt.Println(p)

	// 6. 声明一个函数add，接收两个整数参数并返回它们的和
	add := func(a, b int) int {
		return a + b
	}
	sum := add(2, 3)
	fmt.Println(sum)
}