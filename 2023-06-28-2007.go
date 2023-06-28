package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	dayOfWeek := now.Weekday()

	switch dayOfWeek {
	case time.Monday:
		fmt.Println("今天是星期一")
	case time.Tuesday:
		fmt.Println("今天是星期二")
	case time.Wednesday:
		fmt.Println("今天是星期三")
	case time.Thursday:
		fmt.Println("今天是星期四")
	case time.Friday:
		fmt.Println("今天是星期五")
	case time.Saturday:
		fmt.Println("今天是星期六")
	case time.Sunday:
		fmt.Println("今天是星期日")
	}
}