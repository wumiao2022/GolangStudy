package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成随机数
	randomNumber := rand.Intn(100)

	// 打印随机数
	fmt.Println("随机数为:", randomNumber)

	// 判断随机数的奇偶性
	if randomNumber%2 == 0 {
		fmt.Println("随机数是偶数")
	} else {
		fmt.Println("随机数是奇数")
	}
}