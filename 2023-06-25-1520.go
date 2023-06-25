package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成随机数
	target := rand.Intn(100) + 1
	fmt.Println("猜数字游戏！请输入一个1到100的数字：")

	// 循环获取用户输入并比较
	for {
		var guess int
		fmt.Scan(&guess)

		if guess < 1 || guess > 100 {
			fmt.Println("请输入一个1到100的数字！")
			continue
		}

		if guess > target {
			fmt.Println("猜大了！请重试：")
		} else if guess < target {
			fmt.Println("猜小了！请重试：")
		} else {
			fmt.Println("恭喜你猜对了！")
			break
		}
	}
}