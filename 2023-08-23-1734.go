package main

import (
	"fmt"
)

func main() {
	// 1. 实现一个函数，将一个整数切片翻转
	reverseSlice([]int{1, 2, 3, 4, 5})

	// 2. 编写一个函数，判断一个字符串是否是回文字符串
	isPalindrome("abcba")

	// 3. 实现一个函数，对一个整数切片进行排序（例如使用冒泡排序）
	bubbleSort([]int{5, 3, 2, 4, 1})

	// 4. 编写一个函数，统计一个字符串中每个字符出现的次数
	countChars("hello")

	// 5. 实现一个函数，判断一个整数是否为质数
	isPrime(17)
}

// 1. 实现一个函数，将一个整数切片翻转
func reverseSlice(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	fmt.Println(nums)
}

// 2. 编写一个函数，判断一个字符串是否是回文字符串
func isPalindrome(s string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			fmt.Println("Not palindrome")
			return
		}
	}

	fmt.Println("Palindrome")
}

// 3. 实现一个函数，对一个整数切片进行排序（例如使用冒泡排序）
func bubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	fmt.Println(nums)
}

// 4. 编写一个函数，统计一个字符串中每个字符出现的次数
func countChars(s string) {
	counts := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		counts[s[i]]++
	}

	fmt.Println(counts)
}

// 5. 实现一个函数，判断一个整数是否为质数
func isPrime(num int) {
	if num <= 1 {
		fmt.Println("Not prime")
		return
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			fmt.Println("Not prime")
			return
		}
	}

	fmt.Println("Prime")
}