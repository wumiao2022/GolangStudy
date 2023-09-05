package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 生成随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成随机整数
	randomInt := rand.Intn(100)
	fmt.Println("Generated random integer:", randomInt)

	// 生成随机浮点数
	randomFloat := rand.Float64() * 100
	fmt.Println("Generated random float:", randomFloat)
}