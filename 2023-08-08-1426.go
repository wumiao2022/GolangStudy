package main

import "fmt"

func main() {
	// 练习题1： 打印出从1到100之间所有的奇数
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 练习题2：计算1到100之间所有偶数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)

	// 练习题3：给定一个整数n，判断其是否为素数（质数）
	n := 7
	isPrime := true
	for i := 2; i < n; i++ {
		if n%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(isPrime)
}
