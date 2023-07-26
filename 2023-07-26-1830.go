package main

import "fmt"

func main() {
	// 练习1：打印 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个数的和
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("The sum is:", sum)

	// 练习3：判断一个数的奇偶性
	num := 7
	if num%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	// 练习4：计算圆的面积
	radius := 5.0
	area := 3.14 * radius * radius
	fmt.Println("The area of the circle is:", area)

	// 练习5：判断一个年份是否是闰年
	year := 2020
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Println("It is a leap year")
	} else {
		fmt.Println("It is not a leap year")
	}
}
