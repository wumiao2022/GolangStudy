package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		randomNumber := rand.Intn(100)
		fmt.Printf("Random number %d is %d\n", i+1, randomNumber)
	}
}