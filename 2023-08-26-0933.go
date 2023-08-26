package main

import "fmt"

func main() {
	// 练习1：打印出1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：判断一个数是正数、负数还是零
	num := -5
	if num > 0 {
		fmt.Println("正数")
	} else if num < 0 {
		fmt.Println("负数")
	} else {
		fmt.Println("零")
	}

	// 练习3：计算两个整数的和
	a := 5
	b := 3
	sum := a + b
	fmt.Println("和:", sum)

	// 练习4：计算两个浮点数的乘积
	x := 2.5
	y := 3.7
	product := x * y
	fmt.Println("乘积:", product)

	// 练习5：判断一个年份是否为闰年
	year := 2024
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Println(year, "是闰年")
	} else {
		fmt.Println(year, "不是闰年")
	}
}