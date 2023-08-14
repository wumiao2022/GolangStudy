package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Hello, Golang!")

	// 生成一个随机数
	randomNumber := rand.Intn(100)
	fmt.Println("Generated random number:", randomNumber)

	// 判断随机数的奇偶性
	if randomNumber%2 == 0 {
		fmt.Println("Random number is even")
	} else {
		fmt.Println("Random number is odd")
	}

	// 生成一个随机字符串
	randomString := generateRandomString(10)
	fmt.Println("Generated random string:", randomString)

	// 判断字符串是否包含特定的子串
	if containsSubstring(randomString, "abc") {
		fmt.Println("Random string contains 'abc'")
	} else {
		fmt.Println("Random string does not contain 'abc'")
	}
}

// 生成一个特定长度的随机字符串
func generateRandomString(length int) string {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomString := make([]byte, length)

	for i := range randomString {
		randomString[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(randomString)
}

// 判断字符串是否包含特定的子串
func containsSubstring(str, substr string) bool {
	return !(len(str) < len(substr) || len(substr) <= 0) && str[len(str)-len(substr):] == substr
}