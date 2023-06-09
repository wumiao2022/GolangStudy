package main

import (
	"fmt"
)

func main() {
	// 练习1：打印1~10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：计算1~100所有奇数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	fmt.Println(sum)

	// 练习3：使用星号(*)输出空心的菱形
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			if j == 0 || j == 2*i-2 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	for i := n - 1; i >= 1; i-- {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			if j == 0 || j == 2*i-2 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}