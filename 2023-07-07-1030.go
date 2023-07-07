package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机数
	randomNumber := rand.Intn(100)

	// 判断随机数的奇偶性
	if randomNumber%2 == 0 {
		fmt.Println(randomNumber, "is even")
	} else {
		fmt.Println(randomNumber, "is odd")
	}
}