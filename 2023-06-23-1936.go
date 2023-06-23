package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成随机数
	num := rand.Intn(100)
	fmt.Println(num)

	// 猜数字游戏
	for i := 1; i <= 5; i++ {
		var guess int
		fmt.Print("请输入你猜测的数字：")
		fmt.Scan(&guess)

		if guess == num {
			fmt.Printf("恭喜你猜对了，你用了%d次机会\n", i)
			return
		}

		if guess < num {
			fmt.Println("你猜的数字小了")
		} else {
			fmt.Println("你猜的数字大了")
		}
	}

	fmt.Println("很遗憾，你没有在规定的机会内猜中数字。")
}