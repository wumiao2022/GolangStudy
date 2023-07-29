package main

import (
	"fmt"
)

func main() {
	// 1. 计算两个整数的和并输出
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum =", sum)

	// 2. 打印从1到10的所有奇数
	fmt.Println("Odd numbers from 1 to 10:")
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 3. 判断一个数是否为质数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "is a prime number")
	} else {
		fmt.Println(num, "is not a prime number")
	}

	// 4. 打印一个月的日历
	month := 4
	year := 2022
	daysInMonth := 31
	switch month {
	case 2:
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			daysInMonth = 29
		} else {
			daysInMonth = 28
		}
	case 4, 6, 9, 11:
		daysInMonth = 30
	}
	fmt.Printf("Calendar for %d/%d:\n", month, year)
	fmt.Println("MON TUE WED THU FRI SAT SUN")
	day := 1
	for i := 1; i <= daysInMonth; i++ {
		if i == 1 {
			for j := 1; j < day; j++ {
				fmt.Printf("    ")
			}
		}
		fmt.Printf("%3d ", i)
		if (i+day-1)%7 == 0 || i == daysInMonth {
			fmt.Println()
		}
	}
}