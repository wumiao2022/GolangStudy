package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 生成随机数
	rand.Seed(int64(100)) // 设置随机数种子
	num := rand.Intn(100) // 生成0-100之间的随机数

	// 猜数游戏
	fmt.Println("猜数游戏开始！")
	for i := 0; i < 5; i++ {
		fmt.Printf("第%d次猜，请输入一个0-100之间的整数：", i+1)
		var guess int
		fmt.Scan(&guess)
		if guess == num {
			fmt.Println("恭喜你，猜对了！")
			return
		} else if guess < num {
			fmt.Println("猜小了！")
		} else {
			fmt.Println("猜大了！")
		}
	}
	fmt.Println("很遗憾，你没有猜对！")
	fmt.Printf("正确答案是%d\n", num)
}