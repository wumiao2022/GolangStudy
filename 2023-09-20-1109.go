package main

import "fmt"

func main() {
	// 打印Hello, World!
	fmt.Println("Hello, World!")

	// 计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

	// 判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
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

	// 交换两个变量的值
	a, b := 10, 20
	fmt.Printf("Before swapping: a = %d, b = %d\n", a, b)
	a, b = b, a
	fmt.Printf("After swapping: a = %d, b = %d\n", a, b)
}