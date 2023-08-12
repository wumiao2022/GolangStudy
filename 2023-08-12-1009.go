package main

import (
	"fmt"
)

func main() {
	// 练习1：打印1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	
	// 练习2：计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum of 1 to 100:", sum)
	
	// 练习3：判断一个数是否为奇数
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
	
	// 练习4：输出九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d ", j, i, j*i)
		}
		fmt.Println()
	}
}