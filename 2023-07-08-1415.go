package main

import "fmt"

func main() {
	// Exercise 1
	fmt.Println("Hello, world!")

	// Exercise 2
	fmt.Println(3 + 4)

	// Exercise 3
	name := "John"
	age := 30
	fmt.Printf("%s is %d years old\n", name, age)

	// Exercise 4
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// Exercise 5
	oddNum := 1
	for oddNum <= 10 {
		fmt.Println(oddNum)
		oddNum += 2
	}

	// Exercise 6
	x := 2
	y := 3
	result := multiply(x, y)
	fmt.Printf("%d multiplied by %d is %d\n", x, y, result)
}

func multiply(a, b int) int {
	return a * b
}