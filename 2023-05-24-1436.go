package main

import (
    "fmt"
)

func main() {
    // 练习一：输出 1~100 之间的偶数
    for i := 1; i <= 100; i++ {
        if i%2 == 0 {
            fmt.Print(i, " ")
        }
    }
    fmt.Println()

    // 练习二：使用二分法查找一个有序整数数组中是否存在某个数
    arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
    target := 11
    left := 0
    right := len(arr) - 1
    for left <= right {
        mid := (left + right) / 2
        if arr[mid] == target {
            fmt.Println("Found!")
            return
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    fmt.Println("Not found.")

    // 练习三：实现一个简单的计算器程序，可以对两个数进行加、减、乘、除
    var num1, num2 float64
    var operator string
    fmt.Printf("Please enter a number: ")
    fmt.Scan(&num1)
    fmt.Printf("Please enter another number: ")
    fmt.Scan(&num2)
    fmt.Printf("Please enter an operator (+,-,*,/): ")
    fmt.Scan(&operator)
    switch operator {
    case "+":
        fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
    case "-":
        fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
    case "*":
        fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)
    case "/":
        if num2 != 0 {
            fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)
        } else {
            fmt.Println("Error: denominator cannot be 0")
        }
    default:
        fmt.Println("Error: invalid operator")
    }
}