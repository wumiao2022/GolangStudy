package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算1到10的数字之和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Printf("%d is a prime number: %t\n", num, isPrime)

	// 练习4：将字符串反转
	word := "Hello, Golang!"
	reversed := ""
	for _, char := range word {
		reversed = string(char) + reversed
	}
	fmt.Println("Reversed:", reversed)
}