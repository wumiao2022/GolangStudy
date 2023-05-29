package main

import (
	"fmt"
)

func main() {
	// 练习1：计算斐波那契数列
	fmt.Println(fibonacci(10))

	// 练习2：判断一个数是否为素数
	fmt.Println(isPrime(37))

	// 练习3：找出一个字符串中出现次数最多的字符
	fmt.Println(mostFrequentChar("hello world"))
}

func fibonacci(n int) []int {
	if n < 2 {
		return []int{0, 1}
	}

	fibonacci := []int{0, 1}

	for i := 2; i < n; i++ {
		fibonacci = append(fibonacci, fibonacci[i-2]+fibonacci[i-1])
	}

	return fibonacci
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func mostFrequentChar(s string) string {
	charMap := make(map[rune]int)

	for _, c := range s {
		if _, ok := charMap[c]; ok {
			charMap[c]++
		} else {
			charMap[c] = 1
		}
	}

	maxCount := 0
	maxChar := ""

	for c, count := range charMap {
		if count > maxCount {
			maxCount = count
			maxChar = string(c)
		}
	}

	return maxChar
}