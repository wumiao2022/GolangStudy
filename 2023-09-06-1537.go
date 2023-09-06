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
	fmt.Println(randomNumber)
}
