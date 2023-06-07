package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		fmt.Printf("数字 %d 的平方根是 %.2f\n", num, sqrt(num))
	}
}

func sqrt(n int) float64 {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 0
	}

	var result float64 = float64(n) / 2
	var lastResult float64
	for i := 0; i < 100; i++ {
		lastResult = result
		result = (result + float64(n)/result) / 2
		if result == lastResult {
			break
		}
	}
	return result
}