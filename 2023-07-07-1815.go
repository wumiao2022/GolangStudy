package main

import (
	"fmt"
	"math"
)

// 练习1: 编写一个函数，计算两个整数的和，并返回结果
func sum(a, b int) int {
	return a + b
}

// 练习2: 实现一个函数，计算一个整数数组的平均值，并返回结果
func average(arr []int) float64 {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return float64(sum) / float64(len(arr))
}

// 练习3: 编写一个函数，判断一个数是否为素数，并返回结果
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// 练习4: 编写一个函数，实现字符串的反转，并返回结果
func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	// 练习1: 测试sum函数
	fmt.Println(sum(3, 4)) // 输出: 7

	// 练习2: 测试average函数
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(average(arr)) // 输出: 3

	// 练习3: 测试isPrime函数
	num := 7
	fmt.Println(isPrime(num)) // 输出: true

	// 练习4: 测试reverseString函数
	str := "Hello, World!"
	fmt.Println(reverseString(str)) // 输出: "!dlroW ,olleH"
}