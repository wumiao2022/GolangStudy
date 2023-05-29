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

	// 循环3次，最多猜3次
	for i := 0; i < 3; i++ {
		fmt.Print("请输入一个0到100之间的数字：")
		var guess int
		fmt.Scanln(&guess)

		if guess > num {
			fmt.Println("太大了")
		} else if guess < num {
			fmt.Println("太小了")
		} else {
			fmt.Println("恭喜你猜对了！")
			return
		}
	}

	fmt.Println("很遗憾，你已经猜了3次还没猜中，正确答案是：", num)
}