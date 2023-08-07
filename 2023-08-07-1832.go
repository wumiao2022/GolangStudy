package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：求两个整数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：求两个浮点数的乘积
	float1 := 2.5
	float2 := 3.2
	product := float1 * float2
	fmt.Println("Product:", product)

	// 练习4：判断一个数是否为偶数
	num := 26
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}