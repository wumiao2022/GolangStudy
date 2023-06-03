package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 生成一个随机数
	rand.Seed(time.Now().UnixNano())
	targetNum := rand.Intn(100)

	// 初始化猜测次数
	numGuesses := 0

	// 循环执行猜数游戏
	for {
		// 提示用户输入猜测的数字
		var guess int
		fmt.Print("请输入您的猜测数字（0-99）：")
		fmt.Scan(&guess)

		// 比较猜测的数字和目标数字
		if guess < targetNum {
			fmt.Println("太小了，请重试！")
			numGuesses++
		} else if guess > targetNum {
			fmt.Println("太大了，请重试！")
			numGuesses++
		} else {
			numGuesses++
			fmt.Println("恭喜您，猜对了！您一共猜测了", numGuesses, "次。")
			break
		}
	}
}