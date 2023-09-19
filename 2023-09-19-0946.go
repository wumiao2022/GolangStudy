package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 练习1：生成随机整数
	fmt.Println("随机整数:", rand.Intn(100))

	// 练习2：生成随机浮点数
	fmt.Println("随机浮点数:", rand.Float64())

	// 练习3：生成随机字符串
	fmt.Println("随机字符串:", randomString(10))
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
