package main

import (
	"fmt"
)

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 计算两数之和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3: 判断一个数是奇数还是偶数
	number := 15
	if number%2 == 0 {
		fmt.Printf("%d is even.\n", number)
	} else {
		fmt.Printf("%d is odd.\n", number)
	}

	// 练习4: 判断字符串是否是回文串
	str1 := "level"
	str2 := "hello"
	isPalindrome1 := isPalindrome(str1)
	isPalindrome2 := isPalindrome(str2)
	fmt.Printf("%s is palindrome: %v\n", str1, isPalindrome1)
	fmt.Printf("%s is palindrome: %v\n", str2, isPalindrome2)

	// 练习5: 找出数组中的最大数
	numbers := []int{10, 5, 8, 12, 3, 15}
	max := findMax(numbers)
	fmt.Println("Max number:", max)
}

func isPalindrome(str string) bool {
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}

func findMax(numbers []int) int {
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}
