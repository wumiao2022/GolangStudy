package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机整数
	num := rand.Intn(100)

	// 判断该数是否为奇数
	if num%2 == 0 {
		fmt.Printf("%d 是一个偶数\n", num)
	} else {
		fmt.Printf("%d 是一个奇数\n", num)
	}
}