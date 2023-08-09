package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 设置随机种子

	// 生成一个随机整数
	randomNum := rand.Intn(100) // 生成0~99之间的整数
	fmt.Println("随机整数:", randomNum)

	// 生成一个随机浮点数
	randomFloat := rand.Float64()
	fmt.Println("随机浮点数:", randomFloat)

	// 生成一个随机布尔值
	randomBool := rand.Intn(2) == 0
	fmt.Println("随机布尔值:", randomBool)

	// 生成一个随机字符
	randomChar := rune(rand.Intn(26) + 'a') // 生成随机小写字母 ('a'的unicode码为97)
	fmt.Println("随机字符:", string(randomChar))
}