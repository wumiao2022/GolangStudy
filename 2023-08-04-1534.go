package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	
	// Exercise 1: Generate a random number between 1 and 100
	randomNum := rand.Intn(100) + 1
	fmt.Println(randomNum)

	// Exercise 2: Calculate the square of a random number between 1 and 10
	randomNum2 := rand.Intn(10) + 1
	square := randomNum2 * randomNum2
	fmt.Println(square)

	// Exercise 3: Generate a random string with length 5
	randomString := generateRandomString(5)
	fmt.Println(randomString)
}

func generateRandomString(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
