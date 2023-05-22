package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 生成随机数
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(100) + 1 // 生成1-100之间的随机数

	// 猜数字游戏
	var guess int
	for {
		fmt.Print("请输入一个1-100之间的整数：")
		fmt.Scan(&guess)
		if guess > number {
			fmt.Println("猜大了，请再猜一次")
		} else if guess < number {
			fmt.Println("猜小了，请再猜一次")
		} else {
			fmt.Println("恭喜你，猜对了！")
			break
		}
	}
}