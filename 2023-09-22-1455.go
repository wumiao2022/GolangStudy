package main

import "fmt"

func main() {
	// Exercise 1: Declare two integer variables, assign values of 10 and 20 respectively, and swap their values without using a temporary variable.

	a := 10
	b := 20

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("After swapping, a =", a)
	fmt.Println("After swapping, b =", b)

	// Exercise 2: Write a function to find the factorial of a given integer.

	num := 5
	factorial := 1

	for i := 1; i <= num; i++ {
		factorial *= i
	}

	fmt.Println("Factorial of", num, "is", factorial)

	// Exercise 3: Write a program that prints the Fibonacci sequence up to a given number.

	max := 20
	prev := 0
	current := 1

	fmt.Println("Fibonacci sequence up to", max, ":")

	for current <= max {
		fmt.Print(prev, " ")

		temp := prev
		prev = current
		current = temp + current
	}
	fmt.Println()

	// Exercise 4: Write a program to check if a given number is prime or not.

	checkPrime := 23
	isPrime := true

	if checkPrime == 1 {
		isPrime = false
	} else {
		for i := 2; i <= checkPrime/2; i++ {
			if checkPrime%i == 0 {
				isPrime = false
				break
			}
		}
	}

	if isPrime {
		fmt.Println(checkPrime, "is prime")
	} else {
		fmt.Println(checkPrime, "is not prime")
	}
}