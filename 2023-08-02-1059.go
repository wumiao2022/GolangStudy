package main

import "fmt"

func main() {
	// 练习1：将一个整数转换为二进制字符串
	fmt.Println(intToBinaryString(8)) // 输出: 1000

	// 练习2：判断一个字符串是否是回文字符串
	fmt.Println(isPalindrome("level")) // 输出: true

	// 练习3：找出一个整数数组中的最大值和最小值
	arr := []int{5, 2, 8, 3, 1, 9, 4}
	max, min := findMaxAndMin(arr)
	fmt.Println("Max:", max) // 输出: 9
	fmt.Println("Min:", min) // 输出: 1
}

// 练习1：将一个整数转换为二进制字符串
func intToBinaryString(n int) string {
	return fmt.Sprintf("%b", n)
}

// 练习2：判断一个字符串是否是回文字符串
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// 练习3：找出一个整数数组中的最大值和最小值
func findMaxAndMin(arr []int) (int, int) {
	max, min := arr[0], arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}