package main

import (
	"fmt"
)

func main() {
	// 练习1：打印 Hello World
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)

	// 练习3：计算两个浮点数的乘积
	float1 := 10.5
	float2 := 2.5
	product := float1 * float2
	fmt.Printf("The product of %.1f and %.1f is %.1f\n", float1, float2, product)

	// 练习4：判断一个数是偶数还是奇数
	number := 42
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}

	// 练习5：判断一个年份是否为闰年
	year := 2020
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println(year, "is a leap year")
	} else {
		fmt.Println(year, "is not a leap year")
	}
}
