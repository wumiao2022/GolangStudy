package main

import (
	"fmt"
	"strings"
)

func main() {
	// 练习1：将一个字符串转换成大写
	str := "hello world"
	str = strings.ToUpper(str)
	fmt.Println(str)

	// 练习2：判断一个数字是否为偶数，并打印出结果
	num := 5
	if num%2 == 0 {
		fmt.Println("是偶数")
	} else {
		fmt.Println("不是偶数")
	}

	// 练习3：计算两个数的最大公约数
	num1 := 12
	num2 := 18
	gcd := 1
	for i := 1; i <= num1 && i <= num2; i++ {
		if num1%i == 0 && num2%i == 0 {
			gcd = i
		}
	}
	fmt.Printf("最大公约数: %d\n", gcd)

	// 练习4：判断一个字符串是否为回文字符串
	str2 := "level"
	isPalindrome := true
	for i := 0; i < len(str2)/2; i++ {
		if str2[i] != str2[len(str2)-i-1] {
			isPalindrome = false
			break
		}
	}
	fmt.Println("是否为回文字符串:", isPalindrome)
}