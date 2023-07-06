package main

import (
	"fmt"
	"time"
)

func main() {
	// 实现一个倒计时器，从10秒开始倒计时，每秒打印剩余时间，直到倒计时结束
	countdown := 10
	for countdown > 0 {
		fmt.Println(countdown)
		time.Sleep(time.Second)
		countdown--
	}

	fmt.Println("倒计时结束！")
}