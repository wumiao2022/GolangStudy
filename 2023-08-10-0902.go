package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World!")

	// 练习2：计算两个整数的和
	a := 10
	b := 20
	sum := a + b
	fmt.Println("The sum of", a, "and", b, "is", sum)

	// 练习3：判断一个数是否为偶数
	num := 24
	if num%2 == 0 {
		fmt.Println(num, "is an even number")
	} else {
		fmt.Println(num, "is an odd number")
	}

	// 练习4：使用循环计算1到10的累加和
	total := 0
	for i := 1; i <= 10; i++ {
		total += i
	}
	fmt.Println("The sum of numbers from 1 to 10 is", total)

	// 练习5：判断一个年份是否为闰年
	year := 2020
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println(year, "is a leap year")
	} else {
		fmt.Println(year, "is not a leap year")
	}
}