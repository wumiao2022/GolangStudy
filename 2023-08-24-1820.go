package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：打印当前时间
	currentTime := time.Now()
	fmt.Println(currentTime)

	// 练习2：计算两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是奇数还是偶数
	number := 15
	if number%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 练习4：循环打印数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：使用自定义函数输出乘法口诀表
	multiplicationTable()
}

func multiplicationTable() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		fmt.Println()
	}
}