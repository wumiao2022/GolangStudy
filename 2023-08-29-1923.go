package main

import "fmt"

func main() {
    // Day 1
    fmt.Println("Hello, World!")

    // Day 2
    fmt.Println(add(2, 3))

    // Day 3
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println(sum(numbers))

    // Day 4
    greet()

    // Day 5
    fmt.Println(fibonacci(10))
}

func add(a, b int) int {
    return a + b
}

func sum(nums []int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func greet() {
    fmt.Println("Welcome to the Go programming world!")
}

func fibonacci(n int) []int {
    fib := []int{0, 1}

    for i := 2; i < n; i++ {
        fib = append(fib, fib[i-1]+fib[i-2])
    }

    return fib
}