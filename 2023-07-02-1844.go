package main

import "fmt"

func main() {
	// 实现一个函数，将两个整数相加并返回结果
	func add(a, b int) int {
		return a + b
	}

	// 实现一个函数，判断一个字符串是否是回文字符串
	func isPalindrome(str string) bool {
		for i := 0; i < len(str)/2; i++ {
			if str[i] != str[len(str)-1-i] {
				return false
			}
		}
		return true
	}

	// 实现一个函数，返回一个整数数组中的最大值
	func findMax(nums []int) int {
		max := nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] > max {
				max = nums[i]
			}
		}
		return max
	}

	// 实现一个函数，判断一个整数是否为素数
	func isPrime(num int) bool {
		if num < 2 {
			return false
		}
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	fmt.Println(add(2, 3))
	fmt.Println(isPalindrome("racecar"))
	fmt.Println(findMax([]int{3, 7, 5, 9, 1}))
	fmt.Println(isPrime(17))
}