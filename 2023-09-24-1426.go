package main

import (
	"fmt"
)

func main() {
	// 练习1：打印大写字母A到Z
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Printf("%c ", i)
	}

	fmt.Println()

	// 练习2：计算1到100之间所有整数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1到100之间所有整数的和为：", sum)

	// 练习3：判断一个数是奇数还是偶数
	num := 6
	if num%2 == 0 {
		fmt.Println("该数是偶数")
	} else {
		fmt.Println("该数是奇数")
	}
}