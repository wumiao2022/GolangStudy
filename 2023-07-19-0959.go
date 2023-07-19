package main

import "fmt"

func main() {
	// Hello World
	fmt.Println("Hello, World!")

	// Variables
	var a int = 10
	b := 5
	fmt.Println(a + b)

	// Conditional Statements
	num := 7
	if num < 0 {
		fmt.Println("Negative")
	} else if num == 0 {
		fmt.Println("Zero")
	} else {
		fmt.Println("Positive")
	}

	// Loops
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Arrays
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr[2])

	// Slices
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice[3])

	// Functions
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(3, 4))

	// Pointers
	x := 5
	y := &x
	fmt.Println(*y)

	// Structs
	type Person struct {
		name string
		age  int
	}
	p := Person{name: "John", age: 25}
	fmt.Println(p.age)

	// Maps
	m := make(map[string]int)
	m["key"] = 10
	fmt.Println(m["key"])

	// Channels
	ch := make(chan int)
	go func() {
		ch <- 5
	}()
	fmt.Println(<-ch)
}
