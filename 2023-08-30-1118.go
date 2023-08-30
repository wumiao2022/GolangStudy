package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()

	// 打印当前时间
	fmt.Println("当前时间：", now)

	// 获取当前时间的年份
	year := now.Year()

	// 获取当前时间的月份
	month := now.Month()

	// 获取当前时间的日期
	day := now.Day()

	// 打印年份、月份、日期
	fmt.Println("年份：", year)
	fmt.Println("月份：", month)
	fmt.Println("日期：", day)
}
