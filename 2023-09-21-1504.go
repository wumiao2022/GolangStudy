package main

import "fmt"

func main() {
	// 练习1：打印Hello, world!
	fmt.Println("Hello, world!")

	// 练习2：计算1到10的累加和
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
	fmt.Println(num, "is prime:", isPrime)

	// 练习4：反转一个字符串
	str := "Hello, world!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println("Reversed string:", reversedStr)
}