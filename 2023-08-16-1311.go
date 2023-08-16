package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	
	// 练习1：生成一个随机数并打印出来
	num := rand.Intn(100)
	fmt.Println("随机数:", num)

	// 练习2：将字符串反转并打印出来
	str := "Hello, Golang!"
	reversedStr := reverseString(str)
	fmt.Println("反转后的字符串:", reversedStr)

	// 练习3：计算斐波那契数列的第n项
	n := 10
	fib := fibonacci(n)
	fmt.Printf("斐波那契数列的第%d项为:%d\n", n, fib)
}

// 反转字符串
func reverseString(s string) string {
	length := len(s)
	reversed := make([]byte, length)
	
	for i := 0; i < length; i++ {
		reversed[length-i-1] = s[i]
	}
	
	return string(reversed)
}

// 计算斐波那契数列的第n项
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	
	return b
}