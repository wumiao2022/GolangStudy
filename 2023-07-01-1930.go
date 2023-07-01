package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：变量交换
	a := 10
	b := 20
	a, b = b, a
	fmt.Println(a, b)

	// 练习3：判断奇偶数
	num := 6
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}

	// 练习4：判断年份是否为闰年
	year := 2024
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println(year, "是闰年")
	} else {
		fmt.Println(year, "不是闰年")
	}

	// 练习5：计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1到100的和为:", sum)
}