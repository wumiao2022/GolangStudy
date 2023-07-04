package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：将一个字符串反转
	str := "Hello, World!"
	reverseStr := ""
	for i := len(str)-1; i >= 0; i-- {
		reverseStr += string(str[i])
	}
	fmt.Println(reverseStr)

	// 练习2：打印当前时间
	currentTime := time.Now()
	fmt.Println(currentTime)

	// 练习3：计算斐波那契数列的第n个数
	n := 10
	fibN := fibonacci(n)
	fmt.Println(fibN)
}

// 递归实现斐波那契数列
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2) 
}
