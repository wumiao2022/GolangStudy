package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 打印当前时间
	now := time.Now()
	fmt.Println(now)

	// 2. 将当前时间格式化为指定的字符串
	formattedTime := now.Format("2006-01-02 15:04:05")
	fmt.Println(formattedTime)

	// 3. 获取一周后的日期
	oneWeekLater := now.AddDate(0, 0, 7)
	fmt.Println(oneWeekLater)

	// 4. 判断两个时间是否相等
	isEqual := now.Equal(oneWeekLater)
	fmt.Println(isEqual)

	// 5. 计算两个时间的差值
	duration := now.Sub(oneWeekLater)
	fmt.Printf("Duration: %v\n", duration)
}