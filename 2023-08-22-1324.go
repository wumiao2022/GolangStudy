package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：求和函数
	sum := add(2, 3)
	fmt.Println(sum)

	// 练习3：判断素数
	isPrime := checkPrime(7)
	fmt.Println(isPrime)

	// 练习4：反转字符串
	reverseStr := reverse("Hello")
	fmt.Println(reverseStr)
}

// 练习2：求和函数
func add(a, b int) int {
	return a + b
}

// 练习3：判断素数
func checkPrime(num int) bool {
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

// 练习4：反转字符串
func reverse(str string) string {
	result := ""
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result
}