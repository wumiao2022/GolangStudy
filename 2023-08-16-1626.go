package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个字符串切片
	slice := []string{"Golang", "is", "awesome"}

	// 使用 range 遍历切片
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// 使用 time 包获取当前时间
	currentTime := time.Now()
	fmt.Printf("Current time: %s\n", currentTime.Format("2006-01-02 15:04:05"))
}