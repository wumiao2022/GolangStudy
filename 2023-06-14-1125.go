package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 实例1：生成一个随机数，并判断其是否为偶数
	num := rand.Intn(100)
	if num%2 == 0 {
		fmt.Printf("%d 是偶数\n", num)
	} else {
		fmt.Printf("%d 是奇数\n", num)
	}

	// 实例2：使用for循环输出九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%2d ", j, i, i*j)
		}
		fmt.Println()
	}

	// 实例3：判断一个字符串是否为回文字符串
	str := "level"
	isPalindrome := true
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		fmt.Println(str, "是回文字符串")
	} else {
		fmt.Println(str, "不是回文字符串")
	}
}