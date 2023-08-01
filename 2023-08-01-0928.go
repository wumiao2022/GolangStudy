package main

import "fmt"

func main() {
	// Exercise 1: Declare and print variables of different data types
	var (
		num     int     = 10
		flt     float64 = 3.14
		str     string  = "Hello, world!"
		boolean bool    = true
	)

	fmt.Println(num)
	fmt.Println(flt)
	fmt.Println(str)
	fmt.Println(boolean)

	// Exercise 2: Perform Arithmetic Operations
	var a, b int = 10, 5

	sum := a + b
	diff := a - b
	product := a * b
	quotient := a / b

	fmt.Println(sum)
	fmt.Println(diff)
	fmt.Println(product)
	fmt.Println(quotient)

	// Exercise 3: Use if-else condition to compare numbers
	num1, num2 := 10, 20

	if num1 > num2 {
		fmt.Println("Num1 is greater than Num2")
	} else {
		fmt.Println("Num2 is greater than Num1")
	}

	// Exercise 4: Use for loop to print numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Exercise 5: Create a slice and add elements to it
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// Exercise 6: Use switch case to print the corresponding message based on day of the week
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}

	// Exercise 7: Create a function to calculate the factorial of a number
	num3 := 5
	result := factorial(num3)
	fmt.Println(result)

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}