package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	// 生成一个随机整数
	randomInt := rand.Intn(100)

	fmt.Println("猜数字游戏开始！")
	for {
		fmt.Printf("请输入一个0-99的整数：")
		var guess int
		fmt.Scanf("%d", &guess)

		if guess == randomInt {
			fmt.Println("恭喜你猜对了！")
			break
		} else if guess > randomInt {
			fmt.Println("猜的数字太大了，请再试一次")
		} else {
			fmt.Println("猜的数字太小了，请再试一次")
		}
	}
}