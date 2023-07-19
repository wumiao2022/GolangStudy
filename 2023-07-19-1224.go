package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 生成一个随机数
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100)

	// 打印随机数
	fmt.Println("随机数为:", num)

	// 逆序输出数字的每一位
	reverseDigits(num)
}

func reverseDigits(num int) {
	fmt.Print("逆序输出数字的每一位为:")
	for num > 0 {
		// 取余得到最后一位数字
		digit := num % 10
		// 输出最后一位数字
		fmt.Print(digit)
		// 去掉最后一位数字
		num = num / 10
	}
	fmt.Println()
}
