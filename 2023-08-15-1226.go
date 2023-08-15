package main

import (
	"fmt"
)

func main() {
	// 练习题1：打印1到100之间的所有偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习题2：计算斐波那契数列的前20个数字
	fmt.Println(fibonacci(20))

	// 练习题3：统计字符串中每个字符出现的次数
	str := "hello world"
	charCount := make(map[int32]int)
	for _, char := range str {
		charCount[char]++
	}
	fmt.Println(charCount)

	// 练习题4：判断一个数是否是质数
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(isPrime)
}

func fibonacci(n int) []int {
	fib := make([]int, n)
	fib[0] = 0
	fib[1] = 1
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}