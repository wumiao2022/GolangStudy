package main

import "fmt"

func main() {
	// 练习1: 打印所有小于100的偶数
	for i := 0; i < 100; i = i + 2 {
		fmt.Println(i)
	}

	// 练习2: 计算1到10的累加和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习3: 判断一个数是否是质数
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
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

	// 练习4: 求一个字符串的长度
	str := "Hello, Golang!"
	length := 0
	for range str {
		length++
	}
	fmt.Println("Length:", length)

	// 练习5: 反转一个整数
	num = 12345
	reversed := 0
	for num != 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}
	fmt.Println("Reversed:", reversed)
}