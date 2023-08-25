package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算1到10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为素数
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

	// 练习4：将字符串反转
	str := "Hello, Golang!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println("Reversed string:", reversedStr)
}