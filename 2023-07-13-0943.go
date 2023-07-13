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

	// 猜数字游戏
	guessNum := -1
	guessCount := 0

	for guessNum != num {
		fmt.Print("请输入你猜测的数字：")
		fmt.Scanln(&guessNum)

		guessCount++

		if guessNum > num {
			fmt.Println("猜大了！")
		} else if guessNum < num {
			fmt.Println("猜小了！")
		} else {
			fmt.Printf("恭喜你猜对了！答案是%d，共猜了%d次。\n", num, guessCount)
		}
	}
}
