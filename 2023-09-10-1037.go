package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 生成一个随机数并打印出来
	randomNum := rand.Intn(100)
	fmt.Println(randomNum)

	// 创建一个字符串切片，并对切片进行排序
	strSlice := []string{"dog", "cat", "bird", "elephant", "monkey"}
	for i := 0; i < len(strSlice)-1; i++ {
		for j := i + 1; j < len(strSlice); j++ {
			if strSlice[j] < strSlice[i] {
				strSlice[i], strSlice[j] = strSlice[j], strSlice[i]
			}
		}
	}
	fmt.Println(strSlice)

	// 将一个正整数翻转，并打印翻转后的结果
	num := 12345
	reversedNum := 0
	for num != 0 {
		reversedNum = reversedNum*10 + num%10
		num /= 10
	}
	fmt.Println(reversedNum)
}