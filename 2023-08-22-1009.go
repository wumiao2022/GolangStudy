package main

import "fmt"

func main() {
    // Example 1: Hello World
    fmt.Println("Hello, world!")

    // Example 2: Variables
    var x int = 5
    fmt.Println("The value of x is:", x)

    // Example 3: Constants
    const y = 10
    fmt.Println("The value of y is:", y)

    // Example 4: Array
    array := [3]int{1, 2, 3}
    fmt.Println("The elements in the array are:", array)

    // Example 5: Loops
    for i := 0; i < 5; i++ {
        fmt.Println("Loop iteration:", i)
    }
}