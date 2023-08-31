package main

import "fmt"

func main() {
    
    // Example 1
    fmt.Println("Hello, World!")

    // Example 2
    nums := []int{1, 2, 3, 4, 5}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("Sum:", sum)

    // Example 3
    a := 5
    b := 10
    if a > b {
        fmt.Println("a is greater than b")
    } else if a < b {
        fmt.Println("a is less than b")
    } else {
        fmt.Println("a and b are equal")
    }
    
    // Example 4
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            fmt.Println(i, "is even")
        } else {
            fmt.Println(i, "is odd")
        }
    }

    // Example 5
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)
    fmt.Println("Hello,", name)
}