package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()

	// 输出当前时间的年、月、日、小时、分钟、秒
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	// 输出当前时间的年、月、日、小时、分钟、秒
	fmt.Printf("当前时间为：%d年%d月%d日 %d时%d分%d秒\n", year, month, day, hour, minute, second)
}
