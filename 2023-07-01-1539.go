package main

import "fmt"

func main() {
	// 练习1: 打印Hello World
	fmt.Println("Hello World")

	// 练习2: 计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum of 1 to 100:", sum)

	// 练习3: 判断一个数是奇数还是偶数
	num := 21
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4: 判断一个字符串是不是回文串
	str := "level"
	isPalindrome := true
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		fmt.Println(str, "is a palindrome")
	} else {
		fmt.Println(str, "is not a palindrome")
	}

	// 练习5: 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
