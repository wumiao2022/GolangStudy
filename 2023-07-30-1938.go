package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 多练习题1：生成 10 个随机数，并计算它们的平均值
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}

	sum := 0
	for _, num := range arr {
		sum += num
	}

	average := float64(sum) / float64(len(arr))
	fmt.Println("随机数的平均值为：", average)

	// 多练习题2：判断一个字符串是否为回文字符串
	str := "level"
	isPalindrome := true

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			isPalindrome = false
			break
		}
	}

	fmt.Println("字符串", str, "是回文字符串:", isPalindrome)

	// 多练习题3：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %2d  ", j, i, i*j)
		}
		fmt.Println()
	}
}