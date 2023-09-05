package main

import "fmt"

func main() {
	// 将字符串反转
	str := "Hello, World!"
	reversed := reverseString(str)
	fmt.Println(reversed)

	// 判断一个数是否为素数
	num := 17
	isPrime := checkPrime(num)
	fmt.Println(isPrime)

	// 求斐波那契数列的第n项
	n := 10
	fib := fibonacci(n)
	fmt.Println(fib)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func checkPrime(num int) bool {
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