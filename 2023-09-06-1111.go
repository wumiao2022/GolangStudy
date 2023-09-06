package main

import "fmt"

func main() {
	// Exercise 1: Hello World
	fmt.Println("Hello, World!")

	// Exercise 2: Variables
	var num1 int = 10
	var num2 float64 = 3.14
	var str string = "Golang"
	fmt.Println(num1, num2, str)

	// Exercise 3: If-else statement
	number := 5
	if number%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// Exercise 4: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Exercise 5: Arrays
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// Exercise 6: Slices
	slice := numbers[1:4]
	fmt.Println(slice)

	// Exercise 7: Maps
	dictionary := make(map[string]int)
	dictionary["one"] = 1
	dictionary["two"] = 2
	dictionary["three"] = 3
	fmt.Println(dictionary)

	// Exercise 8: Functions
	sum := add(2, 3)
	fmt.Println(sum)

	// Exercise 9: Pointers
	myNumber := 10
	changeNumber(&myNumber)
	fmt.Println(myNumber)
}

func add(a, b int) int {
	return a + b
}

func changeNumber(number *int) {
	*number = 20
}
