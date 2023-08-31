package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：判断两个数的大小关系
	num1 := 10
	num2 := 5
	if num1 > num2 {
		fmt.Println("num1大于num2")
	} else if num1 < num2 {
		fmt.Println("num1小于num2")
	} else {
		fmt.Println("num1等于num2")
	}

	// 练习3：计算1-10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("1到10的和为：", sum)

	// 练习4：判断一个数是否是偶数
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}

	// 练习5：判断一个年份是否是闰年
	year := 2024
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println(year, "是闰年")
	} else {
		fmt.Println(year, "不是闰年")
	}
}