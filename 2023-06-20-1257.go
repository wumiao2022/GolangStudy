package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 生成一个随机数
	num := rand.Intn(100)

	// 打印提示信息
	fmt.Println("猜测一个 0 到 100 之间的数字:")

	// 循环猜测数字
	for {
		// 获取用户输入
		var guess int
		fmt.Scanln(&guess)

		// 判断猜测数字是否等于随机数
		if guess == num {
			// 猜对了，结束循环
			fmt.Println("猜对了！")
			break
		} else if guess < num {
			// 猜小了，继续循环
			fmt.Println("猜小了，再试一次：")
		} else if guess > num {
			// 猜大了，继续循环
			fmt.Println("猜大了，再试一次：")
		}
	}
}