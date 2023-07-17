package main

import "fmt"

func main() {
	// 练习1: 打印 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 计算两个整数的和并输出结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", sum)

	// 练习3: 判断一个整数是否为偶数并打印结果
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is an even number.")
	} else {
		fmt.Println(num, "is an odd number.")
	}

	// 练习4: 定义一个字符串变量并打印其长度
	str := "Hello, Golang!"
	fmt.Println("The length of the string is", len(str))

	// 练习5: 使用循环输出1到10之间的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}