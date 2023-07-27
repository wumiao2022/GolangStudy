package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	// 生成一个随机数并打印
	randomNumber := rand.Intn(100)
	fmt.Println("随机数是:", randomNumber)

	// 判断随机数是奇数还是偶数
	if randomNumber%2 == 0 {
		fmt.Println("随机数是偶数")
	} else {
		fmt.Println("随机数是奇数")
	}

	// 生成一个随机字符串并打印
	randomString := generateRandomString(10)
	fmt.Println("随机字符串是:", randomString)
}

// 生成一个指定长度的随机字符串
func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())

	// 字符集合
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// 生成随机字符串
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}