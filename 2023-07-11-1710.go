package main

import (
	"fmt"
)

func main() {
	// 练习1：计算两个整数的和
	num1 := 10
	num2 := 20

	sum := num1 + num2
	fmt.Println(sum)

	// 练习2：打印99乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

	// 练习3：判断一个数是正数、负数还是零
	num := -5

	if num > 0 {
		fmt.Println("正数")
	} else if num < 0 {
		fmt.Println("负数")
	} else {
		fmt.Println("零")
	}

	// 练习4：使用数组存储5个姓名，然后将这些姓名打印出来
	names := [5]string{"Alice", "Bob", "Charlie", "David", "Eve"}

	for _, name := range names {
		fmt.Println(name)
	}
}