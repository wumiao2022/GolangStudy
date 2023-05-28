package main

import (
	"fmt"
)

func main() {
	// 练习1：输出Hello World
	fmt.Println("Hello World")

	// 练习2：输出1~10之间的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：计算1+2+3+...+10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习4：判断一个数是否是偶数
	num := 8
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习5：判断一个字符串是否是回文字符串（正着读和反着读都一样）
	str := "ABCCBA"
	flag := true
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			flag = false
			break
		}
	}
	if flag {
		fmt.Println("回文字符串")
	} else {
		fmt.Println("不是回文字符串")
	}
}