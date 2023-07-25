package main

import "fmt"

func main() {
	// 1. 输出Hello, World!
	fmt.Println("Hello, World!")

	// 2. 定义一个变量age并赋值为30，然后输出变量的值
	age := 30
	fmt.Println(age)

	// 3. 定义一个常量pi，并赋值为3.14，然后输出常量的值
	const pi = 3.14
	fmt.Println(pi)

	// 4. 声明一个切片slice，包含整数1、2、3、4、5，并输出slice的长度和容量
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	// 5. 声明一个空的映射map，键为字符串类型，值为整数类型，并输出map的长度
	myMap := make(map[string]int)
	fmt.Println("Length:", len(myMap))
}