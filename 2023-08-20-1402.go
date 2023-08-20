package main

import (
	"fmt"
)

func main() {
	// 1. 打印Hello World
	fmt.Println("Hello World")

	// 2. 数组的遍历
	// 定义一个数组
	numbers := [5]int{1, 2, 3, 4, 5}

	// 遍历数组并打印每个元素
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// 3. 使用函数计算两个整数的和
	// 定义一个函数，接受两个整数参数，返回它们的和
	add := func(a, b int) int {
		return a + b
	}

	// 调用函数并打印结果
	fmt.Println(add(2, 3))

	// 4. 字符串的拼接
	name := "Alice"
	age := 25

	// 使用字符串拼接操作符将字符串与变量拼接起来
	message := "My name is " + name + " and I am " + fmt.Sprint(age) + " years old."

	// 打印拼接后的字符串
	fmt.Println(message)

	// 5. 使用指针修改变量的值
	x := 5

	// 定义一个函数，接受一个指针参数，修改指针指向的变量的值
	modifyValue := func(ptr *int) {
		*ptr = 10
	}

	// 调用函数并打印变量的值
	modifyValue(&x)
	fmt.Println(x)
}