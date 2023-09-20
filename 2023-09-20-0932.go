package main

import "fmt"

// Exercise1 prints "Hello, World!" to the console
func Exercise1() {
    fmt.Println("Hello, World!")
}

// Exercise2 calculates the sum of two integers
func Exercise2(a, b int) int {
    return a + b
}

// Exercise3 reverses a given string
func Exercise3(s string) string {
    runes := []rune(s)
    n := len(runes)
    for i := 0; i < n/2; i++ {
        runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
    }
    return string(runes)
}

// Exercise4 finds the maximum element in a given slice of integers
func Exercise4(slice []int) int {
    max := slice[0]
    for _, num := range slice {
        if num > max {
            max = num
        }
    }
    return max
}

// Exercise5 checks if a given string is a palindrome
func Exercise5(s string) bool {
    runes := []rune(s)
    n := len(runes)
    for i := 0; i < n/2; i++ {
        if runes[i] != runes[n-1-i] {
            return false
        }
    }
    return true
}

func main() {
    Exercise1()
    fmt.Println(Exercise2(3, 4))
    fmt.Println(Exercise3("Golang"))
    fmt.Println(Exercise4([]int{2, 4, 6, 8, 10}))
    fmt.Println(Exercise5("racecar"))
}