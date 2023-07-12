package main

import "fmt"

func main() {
	// 练习1：打印出"Hello, world!"
	fmt.Println("Hello, world!")

	// 练习2：计算两个整数的和并输出结果
	num1 := 12
	num2 := 23
	sum := num1 + num2
	fmt.Println("The sum is:", sum)

	// 练习3：计算一个整数的平方并输出结果
	number := 5
	square := number * number
	fmt.Println("The square is:", square)

	// 练习4：计算一个圆的面积并输出结果
	radius := 10.0
	pi := 3.14159
	area := pi * radius * radius
	fmt.Println("The area of the circle is:", area)

	// 练习5：将一个字符串反转并输出结果
	str := "Hello, World!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println("The reversed string is:", reversedStr)
}