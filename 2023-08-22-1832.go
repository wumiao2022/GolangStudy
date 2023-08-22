package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机整数
	randomNum := rand.Intn(100)
	fmt.Println("随机整数:", randomNum)

	// 生成一个随机浮点数
	randomFloat := rand.Float64()
	fmt.Println("随机浮点数:", randomFloat)

	// 生成一个随机布尔值
	randomBool := rand.Intn(2) == 0
	fmt.Println("随机布尔值:", randomBool)

	// 生成一个随机字符串
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomStr := ""
	for i := 0; i < 10; i++ {
		randomChar := string(chars[rand.Intn(len(chars))])
		randomStr += randomChar
	}
	fmt.Println("随机字符串:", randomStr)
}