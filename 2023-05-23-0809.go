package main

import (
	"fmt"
)

func main() {
	// 练习1: 输出1-100内的所有偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习2: 计算1-100内所有数字的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习3: 将一个正整数反转
	num := 12345
	reverseNum := 0
	for num != 0 {
		reverseNum = reverseNum*10 + num%10
		num /= 10
	}
	fmt.Println("Reverse Number:", reverseNum)
}