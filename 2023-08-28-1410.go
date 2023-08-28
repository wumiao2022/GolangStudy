package main

import "fmt"

func main() {
	// 1. 输出 "Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 声明一个整数变量 x 并赋值为 10，然后输出 x 的值
	x := 10
	fmt.Println(x)

	// 3. 声明一个浮点数变量 y 并赋值为 3.14，然后输出 y 的值
	y := 3.14
	fmt.Println(y)

	// 4. 声明一个字符串变量 str，并赋值为 "Golang is awesome!"，然后输出 str 的值
	str := "Golang is awesome!"
	fmt.Println(str)

	// 5. 声明一个布尔类型变量 flag，并赋值为 true，然后输出 flag 的值
	flag := true
	fmt.Println(flag)

	// 6. 使用循环输出 1 到 10 的所有奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}

	// 7. 将字符串 "Hello, " 和 "World!" 连接起来，并输出结果
	str1 := "Hello, "
	str2 := "World!"
	fmt.Println(str1 + str2)

	// 8. 声明一个切片变量 nums，包含整数 1 到 5，然后输出切片元素
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	// 9. 声明一个空的映射变量 m，然后将键值对 {1: "one", 2: "two"} 添加到映射中，并输出映射元素
	m := make(map[int]string)
	m[1] = "one"
	m[2] = "two"
	fmt.Println(m)

	// 10. 声明一个函数 add，接受两个整数参数，并返回它们的和。使用该函数计算 3 和 5 的和，并输出结果
	add := func(a, b int) int {
		return a + b
	}
	result := add(3, 5)
	fmt.Println(result)
}