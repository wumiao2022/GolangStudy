package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个1-100之间的随机数
	randomNum := rand.Intn(100) + 1

	fmt.Println("猜数字游戏开始！")
	fmt.Println("请猜一个1-100之间的整数，看你能猜中吗？")

	var guess int
	var tries int

	for {
		fmt.Print("请输入你的猜测：")
		fmt.Scan(&guess)

		tries++

		if guess > randomNum {
			fmt.Println("太大了，请继续猜！")
		} else if guess < randomNum {
			fmt.Println("太小了，请继续猜！")
		} else {
			fmt.Printf("恭喜你，猜对了！你一共猜了%d次。\n", tries)
			break
		}
	}
}