package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		// 生成随机整数
		randInt := rand.Intn(100)

		// 判断奇偶性
		if randInt%2 == 0 {
			fmt.Printf("%d 是偶数\n", randInt)
		} else {
			fmt.Printf("%d 是奇数\n", randInt)
		}
	}
}