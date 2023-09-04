package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 打印当前时间
	currentTime := time.Now()
	fmt.Println(currentTime)

	// 2. 计算两个时间点之间的差值
	firstTime := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	secondTime := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	duration := secondTime.Sub(firstTime)
	fmt.Println(duration)

	// 3. 格式化时间
	currentTimeFormatted := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println(currentTimeFormatted)

	// 4. 获取当前时间的年月日时分秒
	year := currentTime.Year()
	month := currentTime.Month()
	day := currentTime.Day()
	hour := currentTime.Hour()
	minute := currentTime.Minute()
	second := currentTime.Second()
	fmt.Println(year, month, day, hour, minute, second)

	// 5. 判断当前年份是否为闰年
	isLeapYear := (year%4 == 0 && year%100 != 0) || year%400 == 0
	fmt.Println(isLeapYear)

	// 6. 获取当前周几
	weekday := currentTime.Weekday()
	fmt.Println(weekday)

	// 7. 计算两个日期之间相差的天数
	startDate := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	daysDiff := int(endDate.Sub(startDate).Hours() / 24)
	fmt.Println(daysDiff)

	// 8. 获取当前月份的天数
	monthDays := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
	fmt.Println(monthDays)

	// 9. 获取当前时间在一年中的第几天
	dayOfYear := currentTime.YearDay()
	fmt.Println(dayOfYear)

	// 10. 比较两个时间的先后
	isBefore := currentTime.Before(secondTime)
	isAfter := currentTime.After(firstTime)
	fmt.Println(isBefore, isAfter)
}