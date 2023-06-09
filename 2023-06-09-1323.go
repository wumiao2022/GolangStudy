package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100)
	fmt.Println("Guess the number!")
	for {
		var guess int
		fmt.Scan(&guess)
		if guess < target {
			fmt.Println("Too low")
		} else if guess > target {
			fmt.Println("Too high")
		} else {
			fmt.Println("Congratulations, you guessed the number!")
			break
		}
	}

}