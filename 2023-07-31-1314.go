package main

import "fmt"

func main() {
	// 练习1
	var name string = "Alice"
	var age int = 25
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// 练习2
	var num1, num2 int = 10, 5
	fmt.Printf("The sum of %d and %d is %d.\n", num1, num2, num1+num2)

	// 练习3
	var isStudent bool = true
	if isStudent {
		fmt.Println("I am a student.")
	} else {
		fmt.Println("I am not a student.")
	}

	// 练习4
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 练习5
	var x int = 5
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x > 5 {
		fmt.Println("x is greater than 5 but less than or equal to 10")
	} else {
		fmt.Println("x is less than or equal to 5")
	}
}