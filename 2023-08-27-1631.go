package main

import "fmt"

func main() {
	// 练习1：输出 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：循环打印数字 1 到 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "is a prime number")
	} else {
		fmt.Println(num, "is not a prime number")
	}

	// 练习4：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

	// 练习5：判断一个字符串是否为回文字符串
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
}
