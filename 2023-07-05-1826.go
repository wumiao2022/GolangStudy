package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数的和
	a := 5
	b := 3
	sum := a + b
	fmt.Println("The sum is:", sum)

	// 练习3：判断一个数是正数、负数还是零
	num := -12
	if num > 0 {
		fmt.Println("The number is positive.")
	} else if num < 0 {
		fmt.Println("The number is negative.")
	} else {
		fmt.Println("The number is zero.")
	}

	// 练习4：计算两个浮点数的乘积
	c := 3.14
	d := 2.5
	product := c * d
	fmt.Println("The product is:", product)

	// 练习5：判断一个年份是否为闰年
	year := 2024
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Println("The year is a leap year.")
	} else {
		fmt.Println("The year is not a leap year.")
	}
}
