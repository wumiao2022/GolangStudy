package main

import "fmt"

func main() {
    // Exercise 1: Variables
    var x int = 5
    var y float64 = 10.5
    sum := x + int(y)
    fmt.Println(sum)

    // Exercise 2: Loops
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // Exercise 3: Functions
    fmt.Println(add(3, 4))
}

func add(a, b int) int {
    return a + b
}