package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()

	// 格式化日期输出
	fmt.Printf("当前时间： %s\n", now.Format("2006-01-02"))

	// 获取当前小时
	hour := now.Hour()

	// 根据小时打印不同的问候语
	if hour < 12 {
		fmt.Println("早上好！")
	} else if hour < 18 {
		fmt.Println("下午好！")
	} else {
		fmt.Println("晚上好！")
	}
}