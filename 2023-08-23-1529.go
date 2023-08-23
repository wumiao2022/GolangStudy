package main

import "fmt"

func main() {
    // 1. 输出Hello, World!
    fmt.Println("Hello, World!")

    // 2. 计算两个整数的和
    a := 10
    b := 20
    sum := a + b
    fmt.Println("Sum:", sum)

    // 3. 判断一个数是否为偶数
    num := 26
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 4. 计算一个数的阶乘
    fact := 1
    for i := 1; i <= num; i++ {
        fact *= i
    }
    fmt.Println("Factorial of", num, "is", fact)

    // 5. 输出斐波那契数列前n个数
    n := 10
    fmt.Print("Fibonacci Series:")
    a, b = 0, 1
    for i := 0; i < n; i++ {
        fmt.Print(a, " ")
        a, b = b, a+b
    }
    fmt.Println()
}