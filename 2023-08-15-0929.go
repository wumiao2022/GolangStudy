package main

import (
	"fmt"
)

func main() {
	// 实现一个求和函数，接收一个整数切片作为参数，返回切片中所有元素的和
	fmt.Println(sum([]int{1, 2, 3, 4, 5})) // 预期输出：15

	// 实现一个去重函数，接收一个整数切片作为参数，返回去重后的切片
	fmt.Println(unique([]int{1, 2, 2, 3, 3, 4, 5, 5})) // 预期输出：[1 2 3 4 5]

	// 实现一个计算阶乘的函数，接收一个非负整数作为参数，返回其阶乘结果
	fmt.Println(factorial(5)) // 预期输出：120

	// 实现一个判断字符串是否为回文的函数，接收一个字符串作为参数，返回布尔值
	fmt.Println(isPalindrome("level")) // 预期输出：true
}

func sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func unique(numbers []int) []int {
	uniqueNumbers := []int{}
	seen := map[int]bool{}
	for _, num := range numbers {
		if !seen[num] {
			uniqueNumbers = append(uniqueNumbers, num)
			seen[num] = true
		}
	}
	return uniqueNumbers
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
```