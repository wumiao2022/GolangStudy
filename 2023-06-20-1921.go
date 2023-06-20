package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子
	for {
		answer := rand.Intn(100) + 1 // 生成1-100的随机数
		fmt.Println("请猜一个1-100之间的数字：")

		for i := 1; i <= 5; i++ { // 最多猜5次
			fmt.Printf("第%d次猜：", i)
			var guess int
			fmt.Scanln(&guess)

			if guess == answer {
				fmt.Println("恭喜你，猜对了！")
				break
			} else if guess < answer {
				fmt.Println("猜小了")
			} else {
				fmt.Println("猜大了")
			}

			if i == 5 {
				fmt.Println("游戏结束，正确答案是：", answer)
			}
		}

		fmt.Println("再来一局吗？（y/n）")
		var playAgain string
		fmt.Scanln(&playAgain)
		if playAgain == "n" {
			fmt.Println("游戏结束，谢谢参与！")
			break
		}
	}
}