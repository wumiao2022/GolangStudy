package main

import (
	"fmt"
)

func main() {
	// 练习1：打印出数字1到10之间的偶数
	fmt.Println("Even numbers between 1 and 10:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习2：计算并打印出数字1到5的累加和
	sum := 0
	fmt.Println("\nSum of numbers from 1 to 5:")
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习3：将一个字符串反转并打印出来
	str := "Hello, World!"
	fmt.Println("\nReversed string:")
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Print(string(str[i]))
	}
	fmt.Println()
}
