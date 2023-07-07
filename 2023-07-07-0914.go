package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, world!")

	// 练习1：打印当前时间
	now := time.Now()
	fmt.Println("Current time:", now.Format("2006-01-02 15:04:05"))

	// 练习2：判断一个数是否为偶数
	num := 10
	if num%2 == 0 {
		fmt.Println(num, "is an even number.")
	} else {
		fmt.Println(num, "is an odd number.")
	}

	// 练习3：计算两个数的和、差、乘积和商
	a := 10
	b := 5
	sum := a + b
	diff := a - b
	product := a * b
	quotient := a / b
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)

	// 练习4：判断一个年份是否为闰年
	year := 2024
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println(year, "is a leap year.")
	} else {
		fmt.Println(year, "is not a leap year.")
	}
}