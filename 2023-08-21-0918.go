package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机整数n（0 <= n <= 100）
	n := rand.Intn(101)

	// 打印出n的平方
	fmt.Println(n * n)
}
