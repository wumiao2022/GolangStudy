package main

import "fmt"

func main() {
    // Exercise 1: Declare and print variables
    var num1, num2 int = 10, 20
    fmt.Println(num1)
    fmt.Println(num2)

    // Exercise 2: Perform arithmetic operations
    sum := num1 + num2
    diff := num1 - num2
    product := num1 * num2
    quotient := num1 / num2
    fmt.Println(sum)
    fmt.Println(diff)
    fmt.Println(product)
    fmt.Println(quotient)

    // Exercise 3: Use if-else statement
    if num1 > num2 {
        fmt.Println("num1 is greater than num2")
    } else {
        fmt.Println("num2 is greater than num1")
    }

    // Exercise 4: Use for loop
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }

    // Exercise 5: Declare and use arrays
    numbers := [5]int{1, 2, 3, 4, 5}
    for i := 0; i < len(numbers); i++ {
        fmt.Println(numbers[i])
    }
}
