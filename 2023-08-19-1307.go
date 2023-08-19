package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机数
	num := rand.Intn(100)

	// 打印随机数
	fmt.Println("随机数:", num)

	// 使用循环判断是否为偶数或奇数
	if num%2 == 0 {
		fmt.Println("这是一个偶数")
	} else {
		fmt.Println("这是一个奇数")
	}
}