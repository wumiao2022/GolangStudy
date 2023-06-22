package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 随机生成一个100以内的整数
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100)
	fmt.Println("随机生成的数字为:", target)

	// 开始猜数字游戏
	fmt.Println("猜数字游戏开始！")
	var guess int
	for {
		fmt.Print("请输入你猜测的数字:")
		fmt.Scanln(&guess)
		if guess < target {
			fmt.Println("猜测的数字偏小，请继续猜测！")
		} else if guess > target {
			fmt.Println("猜测的数字偏大，请继续猜测！")
		} else {
			fmt.Println("恭喜你，猜对了！")
			break
		}
	}
}