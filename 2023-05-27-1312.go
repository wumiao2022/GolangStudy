package main

import (
	"fmt"
)

func main() {
	// 1. 输出9x9乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}

	// 2. 计算斐波那契数列前10项
	fibonacci := []int{0, 1}
	for i := 2; i < 10; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	fmt.Println(fibonacci)

	// 3. 判断一个数是否为质数
	prime := 23
	isPrime := true
	for i := 2; i < prime; i++ {
		if prime%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(prime, "is prime:", isPrime)

	// 4. 判断一个字符串是否为回文字符串
	str := "abcba"
	isPalindrome := true
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			isPalindrome = false
			break
		}
	}
	fmt.Println(str, "is palindrome:", isPalindrome)
}