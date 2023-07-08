package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个数字的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是偶数还是奇数
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：判断一个年份是否是闰年
	year := 2024
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Println(year, "is a leap year")
	} else {
		fmt.Println(year, "is not a leap year")
	}

	// 练习5：遍历切片并求和
	nums := []int{1, 2, 3, 4, 5}
	sum = 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum of nums:", sum)
}