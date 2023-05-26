package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 生成一个 1 到 100 的随机数
	randomNum := rand.Intn(100) + 1

	// 提示玩家猜测数字
	fmt.Println("猜测一个 1 到 100 的数字")

	// 循环猜测，直到玩家猜对了
	for {
		var guess int
		fmt.Scan(&guess)

		if guess == randomNum {
			fmt.Println("恭喜你，猜对了！")
			break
		} else if guess < randomNum {
			fmt.Println("猜的数字太小了，再试试")
		} else {
			fmt.Println("猜的数字太大了，再试试")
		}
	}
}