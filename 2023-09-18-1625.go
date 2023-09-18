package main

import "fmt"

func main() {
	// Exercise 1: Declare two variables, x and y, with initial values of 10 and 5 respectively. Print the sum of x and y.

	x := 10
	y := 5
	sum := x + y
	fmt.Println(sum)

	// Exercise 2: Create a function called multiply that takes two parameters, a and b, and returns their product. Print the result of multiplying 8 and 9 using the multiply function.

	fmt.Println(multiply(8, 9))

	// Exercise 3: Create a function called swapValues that takes two pointers to integers as parameters and swaps their values. Create two integer variables, a and b, with initial values of 10 and 20 respectively. Print the values of a and b before and after calling the swapValues function.

	a := 10
	b := 20
	fmt.Println("Before swapping:")
	fmt.Println("a =", a, "b =", b)
	swapValues(&a, &b)
	fmt.Println("After swapping:")
	fmt.Println("a =", a, "b =", b)

	// Exercise 4: Create a function called calculateArea that takes the radius of a circle as a parameter and returns its area. Create a variable called radius with the value of 5. Print the area of the circle with the given radius using the calculateArea function.

	radius := 5
	fmt.Println(calculateArea(radius))
}

func multiply(a, b int) int {
	return a * b
}

func swapValues(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func calculateArea(radius int) float64 {
	return 3.14 * float64(radius) * float64(radius)
}