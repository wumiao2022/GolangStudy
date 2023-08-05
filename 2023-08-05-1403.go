package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 实现一个函数，生成指定长度的随机字符串
	randomString := generateRandomString(10)
	fmt.Println(randomString)

	// 实现一个函数，计算两个整数的和、差、积、商和取模
	num1 := 10
	num2 := 5
	sum, difference, product, quotient, remainder := calculateOperations(num1, num2)
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func calculateOperations(num1, num2 int) (int, int, int, int, int) {
	sum := num1 + num2
	difference := num1 - num2
	product := num1 * num2
	quotient := num1 / num2
	remainder := num1 % num2
	return sum, difference, product, quotient, remainder
}