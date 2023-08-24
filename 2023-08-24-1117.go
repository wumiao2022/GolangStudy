package main

import "fmt"

// Exercise 1
func exercise1() {
	fmt.Println("Hello, World!")
}

// Exercise 2
func exercise2() {
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Printf("Sum of %d and %d is: %d\n", num1, num2, sum)
}

// Exercise 3
func exercise3() {
	num := 5

	if num%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}
}

// Exercise 4
func exercise4() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// Exercise 5
func exercise5() {
	num := 1

	for num <= 10 {
		fmt.Println(num)
		num++
	}
}

// Exercise 6
func exercise6() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

// Exercise 7
func exercise7() {
	n := 5

	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print("  ")
		}

		for k := 1; k <= i; k++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}
}

// Exercise 8
func exercise8() {
	nums := []int{1, -2, 3, -4, 5}

	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("Sum is:", sum)
}

// Exercise 9
func exercise9() {
	nums := [...]int{1, 2, 3, 4, 5}

	for i, num := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}
}

// Exercise 10
func exercise10() {
	isEven := func(num int) bool {
		if num%2 == 0 {
			return true
		}
		return false
	}

	fmt.Println(isEven(10)) // true
	fmt.Println(isEven(7))  // false
	fmt.Println(isEven(0))  // true
}

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
	exercise7()
	exercise8()
	exercise9()
	exercise10()
}
