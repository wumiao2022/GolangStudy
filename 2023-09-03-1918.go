package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Hello, Go!")
	fmt.Println("Today's date is:", currentTime.Format("2006-01-02"))
	fmt.Println("Let's practice coding in Go!")
}