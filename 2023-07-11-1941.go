package main

import (
	"fmt"
)

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 计算两个整数的和
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Printf("Sum of %d and %d is %d\n", num1, num2, sum)

	// 练习3: 判断一个数是否为奇数
	num := 7
	if num%2 == 0 {
		fmt.Printf("%d is an even number\n", num)
	} else {
		fmt.Printf("%d is an odd number\n", num)
	}

	// 练习4: 根据用户输入的数字，判断是否为质数
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)
	isPrime := true
	for i := 2; i < n; i++ {
		if n%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Printf("%d is a prime number\n", n)
	} else {
		fmt.Printf("%d is not a prime number\n", n)
	}

	// 练习5: 打印斐波那契数列前n项
	n = 10
	fmt.Printf("Fibonacci series for first %d terms:\n", n)
	firstNum := 0
	secondNum := 1
	fmt.Printf("%d\t%d\t", firstNum, secondNum)
	for i := 2; i < n; i++ {
		nextNum := firstNum + secondNum
		fmt.Printf("%d\t", nextNum)
		firstNum = secondNum
		secondNum = nextNum
	}
	fmt.Println()
}