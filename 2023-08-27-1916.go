package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 示例1：生成并打印10个随机整数
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}

	// 示例2：生成并打印10个随机浮点数
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Float64())
	}

	// 示例3：生成并打印一个随机字符串
	fmt.Println(randomString(10))
}

// 生成随机字符串
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}