package main

import "fmt"

func main() {
	// 1. 编写一个函数，判断一个字符串是否是回文字符串，即正反顺序读都是一样的
	// 示例输入："level" -> 输出：true
	// 示例输入："hello" -> 输出：false
	isPalindrome := func(s string) bool {
		for i := 0; i < len(s)/2; i++ {
			if s[i] != s[len(s)-1-i] {
				return false
			}
		}
		return true
	}
	fmt.Println(isPalindrome("level")) // true
	fmt.Println(isPalindrome("hello")) // false

	// 2. 编写一个函数，求斐波那契数列的第n项（从1开始）
	// 示例输入：6 -> 输出：8
	fibonacci := func(n int) int {
		if n <= 2 {
			return 1
		}
		a, b := 1, 1
		for i := 3; i <= n; i++ {
			c := a + b
			a, b = b, c
		}
		return b
	}
	fmt.Println(fibonacci(6)) // 8

	// 3. 编写一个函数，将给定的整数数组nums中的元素翻转并返回
	// 示例输入：[1, 2, 3, 4, 5] -> 输出：[5, 4, 3, 2, 1]
	reverseArray := func(nums []int) []int {
		n := len(nums)
		for i := 0; i < n/2; i++ {
			nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
		}
		return nums
	}
	fmt.Println(reverseArray([]int{1, 2, 3, 4, 5})) // [5, 4, 3, 2, 1]
}