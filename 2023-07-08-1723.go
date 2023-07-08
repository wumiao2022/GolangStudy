package main

import (
	"fmt"
)

func main() {
	// 练习1：求两个数的和
	a := 10
	b := 5
	sum := a + b
	fmt.Println(sum)

	// 练习2：将字符串转换为整数
	str := "123"
	num := 0
	for _, c := range str {
		num = num*10 + int(c-'0')
	}
	fmt.Println(num)

	// 练习3：判断一个数是否是奇数
	number := 7
	if number%2 != 0 {
		fmt.Println("奇数")
	} else {
		fmt.Println("偶数")
	}

	// 练习4：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
