package main

import "fmt"

func main() {
    // 实现一个函数，可以计算两个整数的和
    func add(a, b int) int {
        return a + b
    }

    // 实现一个函数，可以计算两个整数的差
    func subtract(a, b int) int {
        return a - b
    }

    // 实现一个函数，可以计算两个整数的乘积
    func multiply(a, b int) int {
        return a * b
    }

    // 实现一个函数，可以计算两个整数的商
    func divide(a, b int) int {
        return a / b
    }

    // 调用上述函数并打印结果
    fmt.Println(add(2, 3))       // 输出: 5
    fmt.Println(subtract(5, 2))  // 输出: 3
    fmt.Println(multiply(4, 6))  // 输出: 24
    fmt.Println(divide(10, 2))   // 输出: 5
}