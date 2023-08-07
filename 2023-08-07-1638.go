package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World!")

	// 练习2：计算并打印1到10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum of 1 to 10:", sum)

	// 练习3：打印1到10的偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习4：判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i*i <= num; i++ {
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
}