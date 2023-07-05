package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 随机生成一个整数
	num := rand.Intn(100) + 1

	// 输出生成的整数
	fmt.Println("随机生成的整数:", num)

	// 判断生成的整数是否是奇数
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}

	// 判断生成的整数是否是质数
	if isPrime(num) {
		fmt.Println(num, "是质数")
	} else {
		fmt.Println(num, "不是质数")
	}
}

// 判断一个整数是否是质数
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
