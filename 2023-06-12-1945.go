package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 生成一个随机数
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)

	// 循环猜数字
	for {
		var inputNumber int
		fmt.Print("请输入一个0到100的整数：")
		fmt.Scan(&inputNumber)

		if inputNumber == randomNumber {
			fmt.Println("恭喜你，猜对了！")
			break
		} else if inputNumber < randomNumber {
			fmt.Println("猜小了，再试试")
		} else {
			fmt.Println("猜大了，再试试")
		}
	}
}