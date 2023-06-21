package main

import "fmt"

func main() {
	// 练习1：输出Hello World
	fmt.Println("Hello World")

	// 练习2：输出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：交换两个变量的值
	a := 1
	b := 2
	fmt.Println("交换前：a=", a, " b=", b)
	a, b = b, a
	fmt.Println("交换后：a=", a, " b=", b)

	// 练习4：判断一个年份是否为闰年
	year := 2020
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println(year, "年是闰年")
	} else {
		fmt.Println(year, "年不是闰年")
	}

	// 练习5：求一个数的阶乘
	num := 5
	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}
	fmt.Println(num, "的阶乘为：", result)
}