package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) + 1
	guess := 0
	fmt.Println("猜数字游戏！")
	for guess != randomNumber {
		fmt.Print("请输入你猜的数字：")
		fmt.Scan(&guess)
		if guess > randomNumber {
			fmt.Println("猜大了！")
		} else if guess < randomNumber {
			fmt.Println("猜小了！")
		} else {
			fmt.Println("恭喜你，猜对了！")
		}
	}
}