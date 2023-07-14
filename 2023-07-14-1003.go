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

	// 判断随机数是否是偶数
	if randomNumber%2 == 0 {
		fmt.Println(randomNumber, "是一个偶数")
	} else {
		fmt.Println(randomNumber, "是一个奇数")
	}
}
