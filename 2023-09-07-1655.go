package main

import "fmt"

func main() {
	// 1. 输出"Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 计算两个整数的和并输出结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 3. 打印1到10之间的所有奇数
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 4. 判断一个数是否为素数
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

	// 5. 打印菲波那切数列前20项
	a, b := 0, 1
	fmt.Println(a)
	fmt.Println(b)
	for i := 2; i < 20; i++ {
		c := a + b
		fmt.Println(c)
		a = b
		b = c
	}
}