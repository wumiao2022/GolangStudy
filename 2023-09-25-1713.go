package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, Go!")
	fmt.Println("Today's exercise is to print the current time.")

	currentTime := time.Now()
	fmt.Printf("The current time is: %s\n", currentTime.Format("15:04:05"))
}