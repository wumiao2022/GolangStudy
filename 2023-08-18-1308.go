package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Today's date is:", currentTime.Format("2006-01-02"))

	weekday := currentTime.Weekday()
	fmt.Println("Today is:", weekday)

	month := currentTime.Month()
	fmt.Println("Current month is:", month)

	hour := currentTime.Hour()
	fmt.Println("Current hour is:", hour)

	minute := currentTime.Minute()
	fmt.Println("Current minute is:", minute)

	second := currentTime.Second()
	fmt.Println("Current second is:", second)
}