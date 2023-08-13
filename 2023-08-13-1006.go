package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：打印当前日期和时间
	now := time.Now()
	fmt.Println(now)

	// 练习2：计算一个整数数组中的最大值和最小值
	nums := []int{10, 5, 8, 3, 15, 2}
	min, max := nums[0], nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)

	// 练习3：判断一个字符串是否是回文字符串
	str := "level"
	isPalindrome := true
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			isPalindrome = false
			break
		}
	}
	fmt.Println("Is palindrome:", isPalindrome)
}

// 输出：
// 2021-01-01 15:04:05.000000001 +0800 CST m=+0.000000001
// Min: 2
// Max: 15
// Is palindrome: true