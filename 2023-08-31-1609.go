package main

import "fmt"

func main() {
    // Exercise 1: Print "Hello, World!" to the console
    fmt.Println("Hello, World!")

    // Exercise 2: Create a variable named "name" and assign your name to it. Then print the name to the console.
    name := "John Doe"
    fmt.Println(name)

    // Exercise 3: Create a loop that prints all numbers from 1 to 10
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // Exercise 4: Create a function called "add" that takes two integers as parameters and returns their sum.
    fmt.Println(add(5, 3))

    // Exercise 5: Create a struct type called "person" with "name" and "age" fields. Create an instance of the struct and print its values.
    type person struct {
        name string
        age  int
    }
    p := person{"Alice", 25}
    fmt.Println(p.name, p.age)
}

func add(a, b int) int {
    return a + b
}