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
	num := rand.Intn(100)

	// 打印随机数
	fmt.Println(num)

	// 循环5次，每次生成一个随机数并打印
	for i := 0; i < 5; i++ {
		num = rand.Intn(100)
		fmt.Println(num)
	}
}