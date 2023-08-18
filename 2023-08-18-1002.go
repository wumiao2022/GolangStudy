package main

import "fmt"

func main() {
	// 练习1：打印Hello World!
	fmt.Println("Hello World!")

	// 练习2：求1到10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为质数
	num := 29
	prime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			prime = false
			break
		}
	}
	if prime {
		fmt.Println(num, "is prime")
	} else {
		fmt.Println(num, "is not prime")
	}
}