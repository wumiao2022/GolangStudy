package main

import (
    "fmt"
)

func main() {
    // 1. Print "Hello, World!"
    fmt.Println("Hello, World!")

    // 2. Calculate the sum of two numbers and print the result
    a := 5
    b := 10
    sum := a + b
    fmt.Println("Sum:", sum)

    // 3. Create a function to calculate the factorial of a number and print the result
    num := 5
    result := factorial(num)
    fmt.Println("Factorial of", num, ":", result)

    // 4. Create a struct to represent a person with name and age, and print the details
    type Person struct {
        Name string
        Age  int
    }
    person := Person{Name: "John", Age: 30}
    fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
