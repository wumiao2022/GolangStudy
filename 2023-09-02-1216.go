package main

import "fmt"

func main() {
	// 1. 创建一个字符串变量，将其初始值设为"Hello, Go!"
	s := "Hello, Go!"
	fmt.Println(s)

	// 2. 创建一个整数变量，将其初始值设为10
	num := 10
	fmt.Println(num)

	// 3. 创建一个浮点数变量，将其初始值设为3.14
	f := 3.14
	fmt.Println(f)

	// 4. 创建一个布尔变量，将其初始值设为true
	isTrue := true
	fmt.Println(isTrue)

	// 5. 创建一个切片，包含元素1、2、3、4、5
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// 6. 创建一个映射，将键值对"apple"和"red"添加到映射中
	m := make(map[string]string)
	m["apple"] = "red"
	fmt.Println(m)

	// 7. 创建一个函数，将两个整数相加并返回结果
	add := func(a, b int) int {
		return a + b
	}
	result := add(5, 3)
	fmt.Println(result)
}