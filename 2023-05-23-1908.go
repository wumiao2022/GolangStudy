package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100) + 1
	tries := 0 // 记录猜测次数

	fmt.Println("电脑已经想好了一个1-100之间的数字，请你猜一猜是多少？")

	for {
		tries++
		var guess int
		fmt.Printf("请输入你猜测的数字(1-100)：")
		fmt.Scan(&guess)

		if guess < num {
			fmt.Println("猜小了，再试试。")
		} else if guess > num {
			fmt.Println("猜大了，再试试。")
		} else {
			fmt.Printf("恭喜你，猜对了！你一共猜测了%d次。\n", tries)
			break
		}
	}
}