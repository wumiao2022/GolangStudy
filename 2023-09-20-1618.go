package main

import "fmt"

func main() {
    // 1. Hello World
    fmt.Println("Hello, World!")

    // 2. Variables and Constants
    var a = 5
    b := 10
    const c = 15

    // 3. If-else statement
    if a > b {
        fmt.Println("a is greater than b")
    } else {
        fmt.Println("b is greater than a")
    }

    // 4. For loop
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // 5. Arrays and Slices
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println(nums[2])

    // 6. Functions
    res := multiply(2, 3)
    fmt.Println(res)
}

func multiply(a, b int) int {
    return a * b
}