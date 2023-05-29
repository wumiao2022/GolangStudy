package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // 设置随机数种子
	target := rand.Intn(100)     // 随机生成一个目标数字

	for {
		var guess int
		fmt.Print("猜一个1到100的数字：")
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("请输入一个有效的数字！")
		} else {
			if guess < target {
				fmt.Println("猜小了！")
			} else if guess > target {
				fmt.Println("猜大了！")
			} else {
				fmt.Println("猜对了！")
				break
			}
		}
	}
}