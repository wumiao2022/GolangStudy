package main

import "fmt"

func main() {
	// 1. 输出Hello, World!
	fmt.Println("Hello, World!")

	// 2. 声明和初始化一个整数变量并打印
	num := 42
	fmt.Println(num)

	// 3. 声明一个字符串变量并赋值为"Go语言"
	str := "Go语言"
	fmt.Println(str)

	// 4. 计算并打印1+2+3+...+10的结果
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 5. 声明一个切片，包含元素1, 2, 3, 4, 5，并打印切片的长度和容量
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	// 6. 声明一个map，键为字符串类型，值为整数类型，并插入一些键值对，并打印map中的所有键值对
	myMap := make(map[string]int)
	myMap["a"] = 1
	myMap["b"] = 2
	myMap["c"] = 3
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	// 7. 声明一个函数，将两个整数相加并返回结果，然后调用该函数并打印结果
	add := func(a, b int) int {
		return a + b
	}
	result := add(3, 4)
	fmt.Println(result)

	// 8. 声明一个结构体类型，包含姓名和年龄字段，并实例化一个结构体对象并打印其值
	type Person struct {
		Name string
		Age  int
	}
	person := Person{
		Name: "John",
		Age:  30,
	}
	fmt.Println(person)
}