package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机数
	randomNum := rand.Intn(100)

	// 将随机数输出
	fmt.Println(randomNum)
}