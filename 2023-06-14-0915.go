package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成随机数
	randomNum := rand.Intn(100)

	// 提示用户输入数字
	fmt.Println("猜猜看我想的是哪个数字（0-99）：")

	var guessNum int

	count := 0
	for {
		// 读取用户输入
		fmt.Scanln(&guessNum)

		count++

		// 判断用户猜测的数字和随机数的关系
		if guessNum == randomNum {
			fmt.Printf("恭喜你，猜对了！你猜了%d次\n", count)
			break
		} else if guessNum > randomNum {
			fmt.Printf("猜大了，再试试吧（已猜%d次）：\n", count)
		} else {
			fmt.Printf("猜小了，再试试吧（已猜%d次）：\n", count)
		}
	}
}