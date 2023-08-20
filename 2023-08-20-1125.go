package main

import "fmt"

func main() {
    // Example 1: Print "Hello, World!"
    fmt.Println("Hello, World!")

    // Example 2: Print the sum of two numbers (3 and 4)
    sum := 3 + 4
    fmt.Println("Sum:", sum)

    // Example 3: Calculate the factorial of a number (5)
    number := 5
    factorial := 1
    for i := 1; i <= number; i++ {
        factorial *= i
    }
    fmt.Println("Factorial of", number, "is", factorial)

    // Example 4: Reverse a string ("Golang")
    str := "Golang"
    reversedStr := ""
    for i := len(str) - 1; i >= 0; i-- {
        reversedStr += string(str[i])
    }
    fmt.Println("Reversed string:", reversedStr)

    // Example 5: Find the maximum element in an array ([2, 5, 1, 9, 4])
    numbers := []int{2, 5, 1, 9, 4}
    max := numbers[0]
    for _, num := range numbers {
        if num > max {
            max = num
        }
    }
    fmt.Println("Maximum element:", max)
}
