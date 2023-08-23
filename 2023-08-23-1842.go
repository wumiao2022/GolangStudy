package main

import "fmt"

func main() {
	// 实现一个函数，将两个字符串连接起来
	str1 := "Hello"
	str2 := "World"
	result := concatStrings(str1, str2)
	fmt.Println(result)

	// 实现一个函数，将一个整数数组反转
	nums := []int{1, 2, 3, 4, 5}
	reversedNums := reverseArray(nums)
	fmt.Println(reversedNums)

	// 实现一个函数，计算一个字符串中字母的个数
	str := "Hello, World!"
	count := countLetters(str)
	fmt.Println(count)
}

func concatStrings(str1, str2 string) string {
	return str1 + str2
}

func reverseArray(nums []int) []int {
	reversed := make([]int, len(nums))
	for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 {
		reversed[i] = nums[j]
	}
	return reversed
}

func countLetters(str string) int {
	count := 0
	for _, char := range str {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
			count++
		}
	}
	return count
}