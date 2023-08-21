package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, Golang! Let's practice coding!")

	// 实例1：打印当前时间
	fmt.Println("==== Practice 1: Print Current Time ====")
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)

	// 实例2：计算两个数的和
	fmt.Println("==== Practice 2: Calculate Sum ====")
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Printf("%d + %d = %d\n", num1, num2, sum)

	// 实例3：打印指定范围内的奇数
	fmt.Println("==== Practice 3: Print Odd Numbers ====")
	rangeStart := 1
	rangeEnd := 20
	fmt.Printf("Odd numbers between %d and %d:\n", rangeStart, rangeEnd)
	for i := rangeStart; i <= rangeEnd; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 实例4：检查一个数是否为质数
	fmt.Println("==== Practice 4: Check Prime Number ====")
	number := 13
	isPrime := true
	for i := 2; i <= int(number/2); i++ {
		if number%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Printf("%d is a prime number\n", number)
	} else {
		fmt.Printf("%d is not a prime number\n", number)
	}

	fmt.Println("Keep practicing Golang and happy coding!")
}