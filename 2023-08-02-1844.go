package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机数
	randomNumber := rand.Intn(100) + 1

	// 打印随机数
	fmt.Printf("Random number: %d\n", randomNumber)

	// 判断随机数是否为偶数
	if randomNumber%2 == 0 {
		fmt.Printf("The random number %d is even.\n", randomNumber)
	} else {
		fmt.Printf("The random number %d is odd.\n", randomNumber)
	}
}