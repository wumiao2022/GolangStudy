package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()

	fmt.Println("Today is", currentTime.Weekday().String())
	fmt.Println("Current time is", currentTime.Format("15:04:05"))

	for i := 1; i <= 5; i++ {
		fmt.Println("Hello, Golang! This is practice exercise", i)
	}
}