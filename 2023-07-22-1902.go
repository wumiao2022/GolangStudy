package main

import "fmt"

func main() {
	// 练习1: 打印出1到10之间的所有偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习2: 计算并打印出1到100之间所有奇数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	fmt.Println(sum)

	// 练习3: 判断一个字符串是否是回文字符串，即正序和逆序读都一样
	str := "level"
	isPalindrome := true
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			isPalindrome = false
			break
		}
	}
	fmt.Println(isPalindrome)
}