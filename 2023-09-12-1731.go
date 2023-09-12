package main

import "fmt"

func main(){
    // Example 1: Hello World
    fmt.Println("Hello, World!")

    // Example 2: Variables
    var name string = "John"
    age := 30
    fmt.Println("My name is", name, "and I am", age, "years old.")

    // Example 3: Arrays
    var fruits [3]string
    fruits[0] = "Apple"
    fruits[1] = "Banana"
    fruits[2] = "Orange"
    fmt.Println("My favorite fruits are", fruits[0], ",", fruits[1], ", and", fruits[2])

    // Example 4: Slices
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println("The length of numbers slice is", len(numbers))

    // Example 5: Loops
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }

    // Example 6: Functions
    result := multiply(3, 4)
    fmt.Println("3 multiplied by 4 is equal to", result)
}

func multiply(a, b int) int {
    return a * b
}