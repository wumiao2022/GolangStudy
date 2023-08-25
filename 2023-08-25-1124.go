package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 使用当前时间作为随机数种子

	// 生成10个随机数并打印出来
	for i := 0; i < 10; i++ {
		num := rand.Intn(100) // 生成0到99之间的随机数
		fmt.Println(num)
	}
}