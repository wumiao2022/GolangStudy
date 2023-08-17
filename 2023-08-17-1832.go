package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机数
	num := rand.Intn(100)

	// 打印随机数
	fmt.Println("随机数:", num)

	// 判断随机数是奇数还是偶数
	if num%2 == 0 {
		fmt.Println("随机数是偶数")
	} else {
		fmt.Println("随机数是奇数")
	}

	// 判断随机数是否是素数
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println("随机数是素数")
	} else {
		fmt.Println("随机数不是素数")
	}
}