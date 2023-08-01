package main

import "fmt"

// 1. 打印1到10之间的所有整数
func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// 2. 打印一个给定整数的所有因子
func printFactors(num int) {
	fmt.Printf("Factors of %d are: ", num)
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

// 3. 判断一个字符串是否为回文字符串
func isPalindrome(str string) bool {
	// 创建一个反转字符串变量
	var reverseStr string
	// 将字符串反转并赋值给反转字符串变量
	for _, char := range str {
		reverseStr = string(char) + reverseStr
	}
	// 判断原始字符串和反转字符串是否相等
	if str == reverseStr {
		return true
	}
	return false
}

// 4. 计算两个数的最大公约数
func findGCD(a, b int) int {
	// 使用辗转相除法计算最大公约数
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// 程序入口函数
func main() {
	// 测试例子
	printNumbers()

	num := 36
	printFactors(num)

	str := "level"
	fmt.Printf("Is \"%s\" a palindrome? %v\n", str, isPalindrome(str))

	a, b := 24, 36
	fmt.Printf("GCD of %d and %d is: %d\n", a, b, findGCD(a, b))
}