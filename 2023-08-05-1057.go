package main

import "fmt"

// Exercise 1: Hello World
func helloWorld() {
	fmt.Println("Hello, World!")
}

// Exercise 2: Variables
func variables() {
	var name string = "John Doe"
	var age int = 25
	var isStudent bool = true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Student:", isStudent)
}

// Exercise 3: Functions
func add(a, b int) int {
	return a + b
}

// Exercise 4: Loops
func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// Exercise 5: Arrays and Slices
func arraySlices() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Println(names)
}

// Exercise 6: Structs
type Person struct {
	Name string
	Age  int
}

func createPerson() Person {
	return Person{
		Name: "John Doe",
		Age:  25,
	}
}

// Exercise 7: Pointers
func increment(num *int) {
	*num++
}

func main() {
	helloWorld()
	variables()
	fmt.Println("Sum:", add(5, 3))
	printNumbers()
	arraySlices()
	person := createPerson()
	fmt.Println("Person:", person)
	num := 10
	increment(&num)
	fmt.Println("Incremented Number:", num)
}
