package main

import "fmt"

func main() {
	// 练习1
	fmt.Println("Hello, World!")

	// 练习2
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", sum)

	// 练习3
	circleRadius := 3.5
	circleArea := 3.14 * circleRadius * circleRadius
	fmt.Println("The area of the circle with radius", circleRadius, "is", circleArea)

	// 练习4
	numbers := []int{3, 7, 9, 15, 20}
	fmt.Println("Numbers:", numbers)

	// 练习5
	for i, num := range numbers {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}
