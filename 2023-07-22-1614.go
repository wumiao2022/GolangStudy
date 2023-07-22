package main

import (
	"fmt"
)

func main() {
	// 练习1：打印“Hello, World!”
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：计算圆的面积
	const pi = 3.14159
	radius := 5.0
	area := pi * radius * radius
	fmt.Println("Area:", area)

	// 练习4：输出1到10之间的所有奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}

	// 练习5：判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "is a prime number.")
	} else {
		fmt.Println(num, "is not a prime number.")
	}
}