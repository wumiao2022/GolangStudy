package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：输出1到10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为质数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Printf("%d is a prime number\n", num)
	} else {
		fmt.Printf("%d is not a prime number\n", num)
	}

	// 练习4：计算Fibonacci数列前20个数
	first := 0
	second := 1
	fmt.Println("Fibonacci Series:")
	fmt.Print(first, " ", second)
	for i := 3; i <= 20; i++ {
		next := first + second
		fmt.Print(" ", next)
		first = second
		second = next
	}
	fmt.Println()
}