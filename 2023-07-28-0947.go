package main

import "fmt"

func main() {
	// 打印Hello, World!
	fmt.Println("Hello, World!")

	// 整数运算
	x := 10
	y := 5

	// 加法
	fmt.Println("x + y =", x+y)

	// 减法
	fmt.Println("x - y =", x-y)

	// 乘法
	fmt.Println("x * y =", x*y)

	// 除法
	fmt.Println("x / y =", x/y)

	// 取余
	fmt.Println("x % y =", x%y)

	// 字符串操作
	str1 := "Hello"
	str2 := "World"

	// 字符串拼接
	fmt.Println(str1 + ", " + str2)

	// 字符串长度
	fmt.Println("Length of str1 =", len(str1))

	// 条件语句
	if x > y {
		fmt.Println("x is greater than y")
	} else if x < y {
		fmt.Println("x is less than y")
	} else {
		fmt.Println("x is equal to y")
	}

	// 循环语句
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 数组
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	// 切片
	slice := []int{4, 5, 6}
	fmt.Println(slice)

	// 函数
	result := add(3, 4)
	fmt.Println("Result of add function =", result)
}

// 函数定义
func add(a, b int) int {
	return a + b
}