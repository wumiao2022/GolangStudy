package main

import "fmt"

func main() {
	// 练习1：打印十个斐波那契数列的值
	fibonacci(10)

	// 练习2：计算两个整数的最大公约数
	fmt.Println(gcd(18, 24))

	// 练习3：判断一个数是否为素数
	fmt.Println(isPrime(17))

	// 练习4：翻转字符串中的单词顺序
	fmt.Println(reverseWords("Hello World"))

	// 练习5：判断一个字符串是否为回文串
	fmt.Println(isPalindrome("level"))
}

func fibonacci(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func reverseWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}