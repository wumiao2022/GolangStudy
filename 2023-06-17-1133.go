package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var num int = rand.Intn(100) //生成一个0-99的随机数
	var guess int               //输入的猜测数字
	var guessCount int = 0      //猜测次数

	for { //循环直到猜中数字
		fmt.Print("请输入0-99之间的数字：")
		fmt.Scan(&guess)
		guessCount++
		if guess == num {
			fmt.Printf("恭喜你，猜中了数字%d，你一共猜了%d次\n", num, guessCount)
			break
		} else if guess > num {
			fmt.Println("猜测的数字偏大，请继续猜测")
		} else {
			fmt.Println("猜测的数字偏小，请继续猜测")
		}
	}
}