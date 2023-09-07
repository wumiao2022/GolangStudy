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

	// Exercise 3: Arrays
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	// Exercise 4: Slices
	slice := make([]int, 3, 5)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	fmt.Println(slice)

	// Exercise 5: Loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Exercise 6: Functions
	result := sum(2, 3)
	fmt.Println(result)

	// Exercise 7: If-else Condition
	num := 7
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// Exercise 8: Maps
	mymap := make(map[string]int)
	mymap["apple"] = 1
	mymap["banana"] = 2
	mymap["orange"] = 3
	fmt.Println(mymap)

	// Exercise 9: Structs
	type Person struct {
		name string
		age  int
	}
	person := Person{"John", 30}
	fmt.Println(person)

	// Exercise 10: Error Handling
	value, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}

func sum(a, b int) int {
	return a + b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}