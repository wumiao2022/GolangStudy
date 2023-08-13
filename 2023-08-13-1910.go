package main

import (
	"fmt"
	"math"
)

// 练习1：计算两个整数的和
func sum(a, b int) int {
	return a + b
}

// 练习2：计算两个浮点数的乘积
func multiply(a, b float64) float64 {
	return a * b
}

// 练习3：计算一个整数的平方根
func squareRoot(a int) float64 {
	return math.Sqrt(float64(a))
}

// 练习4：判断一个字符串是否是回文字符串
func isPalindrome(str string) bool {
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	// 测试练习1
	fmt.Println(sum(3, 5)) // 输出：8

	// 测试练习2
	fmt.Println(multiply(2.5, 4.2)) // 输出：10.5

	// 测试练习3
	fmt.Printf("%.2f\n", squareRoot(9)) // 输出：3.00

	// 测试练习4
	fmt.Println(isPalindrome("level")) // 输出：true
	fmt.Println(isPalindrome("hello")) // 输出：false
}
```

每日多练，提升自己的Golang编程技能！