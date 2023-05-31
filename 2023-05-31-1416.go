package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	num := rand.Intn(100) + 1
	var guess int
	for {
		fmt.Print("请猜一个1到100的数字：")
		fmt.Scanln(&guess)
		if guess > num {
			fmt.Println("猜的数字太大了，请重试！")
		} else if guess < num {
			fmt.Println("猜的数字太小了，请重试！")
		} else {
			fmt.Println("恭喜你，猜对了！")
			break
		}
	}
}