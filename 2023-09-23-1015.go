package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	currentTime := time.Now()

	// 格式化并输出当前日期
	fmt.Println("Today is:", currentTime.Format("2006-01-02"))

	// 输出当前时间
	fmt.Println("Current time is:", currentTime.Format("15:04:05"))

	// 输出当前年、月、日
	fmt.Println("Year:", currentTime.Year())
	fmt.Println("Month:", currentTime.Month())
	fmt.Println("Day:", currentTime.Day())
}