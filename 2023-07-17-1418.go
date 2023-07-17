package main

import (
	"fmt"
)

func main() {
	// 练习1: 打印1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2: 打印A到Z之间的字母
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Printf("%c ", i)
	}
	fmt.Println()

	// 练习3: 计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习4: 判断一个数是否为素数
	num := 23
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

	// 练习5: 输出斐波那契数列前20项
	fmt.Print("Fibonacci Series:")
	first, second := 0, 1
	for i := 0; i < 20; i++ {
		fmt.Printf("%d ", first)
		temp := first + second
		first = second
		second = temp
	}
	fmt.Println()
}