package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算1+2+...+10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习3：翻转一个字符串
	word := "Hello"
	reversedWord := ""
	for i := len(word) - 1; i >= 0; i-- {
		reversedWord += string(word[i])
	}
	fmt.Println("Reversed word:", reversedWord)

	// 练习4：判断一个数是否为素数
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

	// 练习5：输出斐波那契数列的前n个数
	n := 10
	fibonacci := []int{0, 1}
	for i := 2; i < n; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	fmt.Println("Fibonacci sequence:", fibonacci)
}