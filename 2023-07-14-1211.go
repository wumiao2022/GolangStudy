package main

import "fmt"

// Exercise 1
func exercise1() {
    fmt.Println("Hello, World!")
}

// Exercise 2
func exercise2() {
    var num1, num2 int = 10, 20
    sum := num1 + num2
    fmt.Println("Sum:", sum)
}

// Exercise 3
func exercise3() {
    var radius float64 = 5.0
    const pi float64 = 3.14
    area := pi * radius * radius
    fmt.Println("Area:", area)
}

// Exercise 4
func exercise4() {
    var num int = 18
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }
}

// Exercise 5
func exercise5() {
    var numbers = []int{4, 2, 9, 6, 7}
    for i := 0; i < len(numbers); i++ {
        fmt.Println(numbers[i])
    }
}

func main() {
    exercise1() // Hello, World!
    exercise2() // Sum: 30
    exercise3() // Area: 78.5
    exercise4() // 18 is even
    exercise5() // 4
                // 2
                // 9
                // 6
                // 7
}