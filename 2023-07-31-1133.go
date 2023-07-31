package main

import (
	"fmt"
)

// Exercise1: Hello World
func main() {
	fmt.Println("Hello, World!")
}

// Exercise2: Print Name
func main() {
	fmt.Println("My name is John Doe.")
}

// Exercise3: Sum of Two Numbers
func main() {
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)
}

// Exercise4: Swap Numbers
func main() {
	num1 := 10
	num2 := 20
	temp := num1
	num1 = num2
	num2 = temp
	fmt.Println("Swap Result: num1 =", num1, "num2 =", num2)
}

// Exercise5: Even or Odd
func main() {
	num := 25
	if num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}
}

// Exercise6: Fibonacci Series
func main() {
	num := 10
	a, b := 0, 1
	fmt.Print("Fibonacci Series:", a, ", ", b)
	for i := 2; i < num; i++ {
		next := a + b
		fmt.Print(", ", next)
		a = b
		b = next
	}
	fmt.Println()
}

// Exercise7: Factorial of a Number
func main() {
	num := 5
	fact := 1
	for i := 1; i <= num; i++ {
		fact *= i
	}
	fmt.Println("Factorial of", num, "is", fact)
}

// Exercise8: Array Sum
func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum of array:", sum)
}

// Exercise9: Check Prime Number
func main() {
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "is a prime number.")
	} else {
		fmt.Println(num, "is not a prime number.")
	}
}

// Exercise10: Print Multiplication Table
func main() {
	num := 5
	for i := 1; i <= 10; i++ {
		fmt.Println(num, "*", i, "=", num*i)
	}
}
