package main

import "fmt"

func main() {
	// 实例1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 实例2：计算两个整数的和
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Printf("Sum of %d and %d is %d\n", num1, num2, sum)

	// 实例3：判断某个年份是否是闰年
	year := 2022
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Printf("%d is a leap year\n", year)
	} else {
		fmt.Printf("%d is not a leap year\n", year)
	}

	// 实例4：使用循环打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}