package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机数n，n的取值范围为[1,100]
	n := rand.Intn(100) + 1
	fmt.Printf("我在1-100之间想了一个数，你猜猜看是多少？\n")

	for {
		// 玩家输入一个猜测的数guess
		var guess int
		fmt.Scanln(&guess)

		// 判断玩家的猜测数与随机数n的大小关系，并进行相应的提示
		if guess < n {
			fmt.Printf("%d有点小了，再给你一次机会猜猜看。\n", guess)
		} else if guess > n {
			fmt.Printf("%d有点大了，再给你一次机会猜猜看。\n", guess)
		} else {
			fmt.Printf("你猜中了，这个数就是%d！\n", guess)
			break
		}
	}
}