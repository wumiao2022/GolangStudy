package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(10)

	if randomNum%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}