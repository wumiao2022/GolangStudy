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

	// 打印随机数
	fmt.Println("随机数是：", randomNum)

	// 判断随机数是否是偶数
	if randomNum%2 == 0 {
		fmt.Println("随机数是偶数")
	} else {
		fmt.Println("随机数是奇数")
	}
}