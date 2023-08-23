package main

import "fmt"

func main() {
	a := 5
	b := 10

	sum := add(a, b)
	fmt.Printf("The sum of %d and %d is %d\n", a, b, sum)

	c := 3
	d := 7

	product := multiply(c, d)
	fmt.Printf("The product of %d and %d is %d\n", c, d, product)

	e := 15
	f := 2

	quotient, remainder := divide(e, f)
	fmt.Printf("The quotient and remainder of %d divided by %d are %d and %d respectively\n", e, f, quotient, remainder)

}

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) (int, int) {
	return x / y, x % y
}