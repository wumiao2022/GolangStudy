package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成10个随机数并打印
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}
}