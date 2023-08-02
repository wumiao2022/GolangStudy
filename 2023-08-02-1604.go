package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：判断一个数是否为偶数
	num := 8
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习4：计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1到100的和为:", sum)

	// 练习5：判断一个年份是否为闰年
	year := 2020
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Println(year, "是闰年")
	} else {
		fmt.Println(year, "不是闰年")
	}
}