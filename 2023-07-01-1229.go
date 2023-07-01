package main

import "fmt"

func main() {
    // Exercise 1: Printing "Hello, World!"
    fmt.Println("Hello, World!")

    // Exercise 2: Variables and Constants
    var num1 int = 10
    var num2 float64 = 3.14
    const pi = 3.1416
    fmt.Println(num1)
    fmt.Println(num2)
    fmt.Println(pi)

    // Exercise 3: Data Types and Conversion
    var name string = "John"
    var age int = 25
    fmt.Println("My name is " + name)
    fmt.Println("I am", age, "years old")

    // Exercise 4: Operators
    var x int = 5
    var y int = 2
    fmt.Println("x + y =", x+y)
    fmt.Println("x - y =", x-y)
    fmt.Println("x * y =", x*y)
    fmt.Println("x / y =", x/y)
    fmt.Println("x % y =", x%y)

    // Exercise 5: Conditional Statements
    var a int = 10
    var b int = 5
    if a > b {
        fmt.Println("a is greater than b")
    } else if a < b {
        fmt.Println("a is smaller than b")
    } else {
        fmt.Println("a is equal to b")
    }

    // Exercise 6: Loops
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }

    // Exercise 7: Arrays
    var fruits [3]string
    fruits[0] = "Apple"
    fruits[1] = "Banana"
    fruits[2] = "Orange"
    fmt.Println(fruits)

    // Exercise 8: Slices
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println("First element:", numbers[0])
    fmt.Println("Slice length:", len(numbers))

    // Exercise 9: Functions
    sum := add(2, 3)
    fmt.Println("Sum:", sum)
}

func add(a, b int) int {
    return a + b
}