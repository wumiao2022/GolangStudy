package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：打印1到10之间的所有奇数
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 练习3：计算1到100之间所有偶数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println("Sum of even numbers from 1 to 100:", sum)

	// 练习4：判断一个数是否是质数
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

	// 练习5：求斐波那契数列的前20个数
	fibonacci := []int{0, 1}
	for i := 2; i < 20; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	fmt.Println("Fibonacci Series:", fibonacci)
}