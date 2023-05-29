package main

import (
	"fmt"
)

func main() {
	// 练习1：循环打印数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：判断奇偶数
	num := 5
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习3：判断闰年
	year := 2020
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println("闰年")
	} else {
		fmt.Println("平年")
	}

	// 练习4：二进制转十进制
	binary := "1010"
	decimal := 0
	for i := len(binary) - 1; i >= 0; i-- {
		if binary[i] == '1' {
			decimal += 1 << (len(binary) - 1 - i)
		}
	}
	fmt.Println(decimal)

	// 练习5：闭包函数
	add := func(a int) func(int) int {
		return func(b int) int {
			return a + b
		}
	}
	fmt.Println(add(1)(2))
}