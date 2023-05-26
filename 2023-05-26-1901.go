package main

import (
	"fmt"
	"math"
)

func main() {
	// Example 1: Calculate the area of a circle
	radius := 6.0
	area := math.Pow(radius, 2) * math.Pi
	fmt.Println("Area of circle with radius", radius, "is", area)

	// Example 2: Print the first 10 even numbers using a loop
	count := 0
	for i := 0; count < 10; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
			count++
		}
	}

	// Example 3: Check if a number is prime
	num := 23
	isPrime := true
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println("\n", num, "is a prime number")
	} else {
		fmt.Println("\n", num, "is not a prime number")
	}
}