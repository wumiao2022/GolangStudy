package main

import "fmt"

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 计算两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum =", sum)

	// 练习3: 判断一个数是否为偶数
	num := 16
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4: 使用for循环打印1到10的所有数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5: 定义一个函数，接收一个字符串并返回其长度
	str := "Hello, Golang"
	length := getLength(str)
	fmt.Println("Length =", length)
}

// getLength函数用于计算字符串的长度
func getLength(str string) int {
	return len(str)
}
