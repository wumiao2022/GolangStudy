package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个数的和并打印结果
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：使用循环输出1到10的所有数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4：定义一个函数，接收一个字符串参数并返回字符串的长度
	str := "Hello, Golang!"
	fmt.Println("Length of str:", len(str))

	// 练习5：使用切片将字符串转换为单个字符并逆序打印
	strSlice := []rune(str)
	for i := len(strSlice) - 1; i >= 0; i-- {
		fmt.Print(string(strSlice[i]))
	}
	fmt.Println()
}