package main

import "fmt"

func main() {
	// 练习1: 打印1到100之间的所有偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习2: 计算1到100之间所有奇数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	fmt.Println("Sum of odd numbers between 1 and 100:", sum)

	// 练习3: 反转字符串
	str := "Hello, world!"
	reversed := ""
	for i := len(str)-1; i >= 0; i-- {
		reversed += string(str[i])
	}
	fmt.Println("Reversed string:", reversed)

	// 练习4: 判断一个数是不是质数
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
}