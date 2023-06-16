package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	// 练习1: 生成一个随机数，并判断是奇数还是偶数
	number := rand.Intn(100)
	if number%2 == 0 {
		fmt.Printf("%d 是偶数\n", number)
	} else {
		fmt.Printf("%d 是奇数\n", number)
	}

	// 练习2: 计算出1-100之间的奇数和偶数之和
	sumOdd, sumEven := 0, 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sumEven += i
		} else {
			sumOdd += i
		}
	}
	fmt.Printf("1~100之间的奇数和为%d，偶数和为%d\n", sumOdd, sumEven)

	// 练习3: 定义一个函数，判断一个数是否是质数
	isPrime := func(num int) bool {
		if num <= 1 {
			return false
		}
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}
	number = rand.Intn(100)
	if isPrime(number) {
		fmt.Printf("%d 是质数\n", number)
	} else {
		fmt.Printf("%d 不是质数\n", number)
	}
}