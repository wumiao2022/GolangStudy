package main

import "fmt"

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 计算两个整数的和
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3: 判断一个数是否为偶数
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4: 创建一个字符串数组，并遍历打印数组中的元素
	strArr := [3]string{"apple", "banana", "orange"}
	for _, str := range strArr {
		fmt.Println(str)
	}

	// 练习5: 定义一个函数，接收两个整数参数，并返回它们的乘积
	product := multiply(3, 4)
	fmt.Println("Product:", product)
}

func multiply(a, b int) int {
	return a * b
}
