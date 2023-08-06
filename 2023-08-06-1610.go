package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：打印1到10的平方
	for i := 1; i <= 10; i++ {
		fmt.Println(i*i)
	}

	// 练习3：判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "is a prime number")
	} else {
		fmt.Println(num, "is not a prime number")
	}

	// 练习4：计算斐波那契数列
	n := 10
	fibonacci := make([]int, n)
	fibonacci[0] = 0
	fibonacci[1] = 1
	for i := 2; i < n; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}
	fmt.Println(fibonacci)
}
