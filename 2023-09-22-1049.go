package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：输出当前日期和时间
	currentTime := time.Now()
	fmt.Println(currentTime)

	// 练习2：打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：计算1加到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习4：判断一个数是否为素数
	number := 11
	isPrime := true
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(number, "is prime")
	} else {
		fmt.Println(number, "is not prime")
	}

	// 练习5：翻转一个字符串
	word := "hello"
	reversedWord := ""
	for i := len(word) - 1; i >= 0; i-- {
		reversedWord += string(word[i])
	}
	fmt.Println(reversedWord)
}