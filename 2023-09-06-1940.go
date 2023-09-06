package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 练习1: 生成一个1到100之间的随机整数，并打印出来
	num := rand.Intn(100) + 1
	fmt.Println(num)

	// 练习2: 生成一个长度为10的随机字符串，并打印出来
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	str := make([]rune, 10)
	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}
	fmt.Println(string(str))

	// 练习3: 生成一个长度为5的随机整数切片，并打印出来
	slice := make([]int, 5)
	for i := range slice {
		slice[i] = rand.Intn(100)
	}
	fmt.Println(slice)

	// 练习4: 生成一个长度为5的随机整数数组，并打印出来
	arr := [5]int{}
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
}

注意：这只是一个示例，实际的练习可能因为涉及题目要求的差异而不同。每日多练才能熟能生巧！