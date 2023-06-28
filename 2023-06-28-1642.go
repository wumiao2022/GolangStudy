package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：打印1~10之间的奇数
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 练习4：判断一个数是否为素数
	num := 97
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

	// 练习5：计算斐波那契数列的前20个数字
	num1, num2 = 0, 1
	fmt.Println(num1)
	fmt.Println(num2)
	for i := 3; i <= 20; i++ {
		nextNum := num1 + num2
		fmt.Println(nextNum)
		num1, num2 = num2, nextNum
	}
}
