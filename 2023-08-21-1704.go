package main

import (
	"fmt"
)

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 使用for循环打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3: 使用if-else语句判断一个数是奇数还是偶数
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}

	// 练习4: 使用switch语句判断一个月份有多少天
	month := 2
	year := 2022
	days := 0
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
	case 4, 6, 9, 11:
		days = 30
	case 2:
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			days = 29
		} else {
			days = 28
		}
	default:
		fmt.Println("无效的月份")
	}
	fmt.Printf("%d年%d月有%d天\n", year, month, days)

	// 练习5: 定义一个切片，包含一些整数，使用range循环计算切片中整数的和
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("切片中整数的和为:", sum)
}
```