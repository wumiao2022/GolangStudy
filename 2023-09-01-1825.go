package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：求和函数
	num1 := 10
	num2 := 20
	sum := add(num1, num2)
	fmt.Printf("Sum of %d and %d is %d\n", num1, num2, sum)

	// 练习3：计算数字平方
	num := 5
	square := calcSquare(num)
	fmt.Printf("Square of %d is %d\n", num, square)

	// 练习4：判断字符串是否为回文
	str := "racecar"
	isPalindrome := checkPalindrome(str)
	fmt.Printf("Is %s a palindrome? %t\n", str, isPalindrome)
}

// 求和函数
func add(a, b int) int {
	return a + b
}

// 计算数字平方
func calcSquare(num int) int {
	return num * num
}

// 判断字符串是否为回文
func checkPalindrome(str string) bool {
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return true
}
