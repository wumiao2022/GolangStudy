package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 10; i++ {
		num := rand.Intn(100)
		fmt.Printf("Random number %d: %d\n", i, num)
	}
}