package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")
	
	// 练习2：求和函数
	result := sum(3, 4)
	fmt.Println("3 + 4 =", result)
	
	// 练习3：判断质数
	number := 7
	if isPrime(number) {
		fmt.Println(number, "is a prime number")
	} else {
		fmt.Println(number, "is not a prime number")
	}
}

// 求和函数
func sum(a, b int) int {
	return a + b
}

// 判断质数函数
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
