package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, Golang!")

	// 练习1：打印当前时间
	fmt.Println("当前时间：", time.Now())

	// 练习2：循环打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：判断一个数字是奇数还是偶数
	num := 9
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}

	// 练习4：计算两个数字的和
	num1 := 5
	num2 := 7
	sum := num1 + num2
	fmt.Println(num1, "+", num2, "=", sum)

	// 练习5：使用函数打印九九乘法表
	printMultiplicationTable()

}

func printMultiplicationTable() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		fmt.Println()
	}
}
