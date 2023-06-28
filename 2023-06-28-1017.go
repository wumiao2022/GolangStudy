package main

import (
	"fmt"
)

func main() {
	// 1. 输出"Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 输出1到10之间的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 3. 计算1到100之间所有整数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 4. 判断一个数是否是偶数
	num := 12
	if num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}

	// 5. 计算一个数的阶乘
	num = 5
	factorial := 1
	for i := 1; i <= num; i++ {
		factorial *= i
	}
	fmt.Println("Factorial of", num, "is", factorial)

	// 6. 判断一个字符串是否是回文字符串
	word := "level"
	isPalindrome := true
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-1-i] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		fmt.Println(word, "is a palindrome.")
	} else {
		fmt.Println(word, "is not a palindrome.")
	}
}