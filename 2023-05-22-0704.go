package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成随机数的范围
	min := 1
	max := 100

	// 生成随机数
	answer := rand.Intn(max-min+1) + min

	fmt.Println("猜数字游戏")
	fmt.Printf("请猜一个 %d 到 %d 的数字：", min, max)

	// 循环猜数字
	for {
		var guess int
		fmt.Scanln(&guess)

		// 判断猜测的数字大小
		if guess < answer {
			fmt.Println("猜小了，请重试：")
		} else if guess > answer {
			fmt.Println("猜大了，请重试：")
		} else {
			fmt.Println("恭喜你，猜对了！")
			break
		}
	}
}