package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机数
	randomNumber := rand.Intn(100)
	fmt.Println("随机数:", randomNumber)

	// 使用循环猜测这个数字
	for {
		var guess int
		fmt.Print("请猜测一个数字: ")
		fmt.Scan(&guess)

		if guess == randomNumber {
			fmt.Println("恭喜你，猜对了！")
			break
		} else if guess < randomNumber {
			fmt.Println("很抱歉，猜小了，请继续猜测！")
		} else {
			fmt.Println("很抱歉，猜大了，请继续猜测！")
		}
	}
}