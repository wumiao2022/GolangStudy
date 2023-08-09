package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()

	// 格式化输出当前日期和时间
	fmt.Println("Current time:", now.Format("2006-01-02 15:04:05"))

	// 获取明天的日期
	tomorrow := now.AddDate(0, 0, 1)

	// 格式化输出明天的日期
	fmt.Println("Tomorrow: ", tomorrow.Format("2006-01-02"))
}
