package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：输出整数和浮点数的和
	var num1 int = 10
	var num2 float64 = 3.14
	sum := float64(num1) + num2
	fmt.Println("Sum:", sum)

	// 练习3：输出字符串的长度
	str := "Golang"
	fmt.Println("Length of string:", len(str))

	// 练习4：判断是否是偶数
	num := 8
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习5：计算数组中元素的和
	arr := []int{1, 2, 3, 4, 5}
	sum = 0
	for _, value := range arr {
		sum += value
	}
	fmt.Println("Sum of array elements:", sum)
}