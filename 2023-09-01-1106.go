package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, Go!")
	fmt.Println("Today's exercise is to create a program that calculates the current time.")

	currentTime := time.Now()
	fmt.Println("Current time is:", currentTime.Format(time.RFC3339))
}