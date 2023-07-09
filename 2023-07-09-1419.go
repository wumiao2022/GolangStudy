package main

import "fmt"

func main() {
  // Example 1: Print "Hello, World!"
  fmt.Println("Hello, World!")

  // Example 2: Calculate the sum of two numbers and print the result
  num1 := 10
  num2 := 20
  sum := num1 + num2
  fmt.Println("Sum:", sum)

  // Example 3: Check whether a number is even or odd
  num := 7
  if num%2 == 0 {
    fmt.Println(num, "is even")
  } else {
    fmt.Println(num, "is odd")
  }

  // Example 4: Print numbers from 1 to 10 in a loop
  for i := 1; i <= 10; i++ {
    fmt.Println(i)
  }
}