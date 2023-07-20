package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, Go! Today's exercise is about time.")

	// Exercise 1: Print the current date and time
	currentTime := time.Now()
	fmt.Println("Current date and time:", currentTime)

	// Exercise 2: Get the current year
	currentYear := currentTime.Year()
	fmt.Println("Current year:", currentYear)

	// Exercise 3: Get the current month
	currentMonth := currentTime.Month()
	fmt.Println("Current month:", currentMonth)

	// Exercise 4: Get the current day
	currentDay := currentTime.Day()
	fmt.Println("Current day:", currentDay)

	// Exercise 5: Get the current hour
	currentHour := currentTime.Hour()
	fmt.Println("Current hour:", currentHour)

	// Exercise 6: Get the current minute
	currentMinute := currentTime.Minute()
	fmt.Println("Current minute:", currentMinute)

	// Exercise 7: Get the current second
	currentSecond := currentTime.Second()
	fmt.Println("Current second:", currentSecond)
}