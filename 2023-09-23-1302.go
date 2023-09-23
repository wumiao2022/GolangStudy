package main

import "fmt"

func main() {
    // 练习1
    var num1 int = 10
    var num2 int = 20
    result := addNumbers(num1, num2)
    fmt.Println("Result:", result)

    // 练习2
    name := "Golang"
    printMessage(name)

    // 练习3
    nums := []int{1, 2, 3, 4, 5}
    sum := calculateSum(nums)
    fmt.Println("Sum:", sum)

    // 练习4
    a, b := swapValues(5, 10)
    fmt.Println("Swapped values:", a, b)
}

func addNumbers(a, b int) int {
    return a + b
}

func printMessage(msg string) {
    fmt.Println("Message:", msg)
}

func calculateSum(nums []int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return sum
}

func swapValues(a, b int) (int, int) {
    return b, a
}