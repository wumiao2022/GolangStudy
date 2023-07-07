package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算1到10的和并打印结果
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
	fmt.Println(num, "is prime:", isPrime)

	// 练习4：生成斐波那契数列并打印前10个数
	fib := make([]int, 10)
	fib[0], fib[1] = 0, 1
	for i := 2; i < 10; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	fmt.Println("Fibonacci sequence:", fib)

	// 练习5：将字符串反转并打印结果
	str := "Golang"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println("Reversed string:", reversedStr)
}