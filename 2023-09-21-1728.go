package main

import "fmt"

func main() {
    // 1. 打印 "Hello, World!"
    fmt.Println("Hello, World!")

    // 2. 变量赋值与打印
    var x int = 10
    var y int = 5
    fmt.Println("x =", x)
    fmt.Println("y =", y)

    // 3. 算术运算
    sum := x + y
    difference := x - y
    product := x * y
    quotient := x / y
    fmt.Println("sum =", sum)
    fmt.Println("difference =", difference)
    fmt.Println("product =", product)
    fmt.Println("quotient =", quotient)

    // 4. 条件语句
    if x > y {
        fmt.Println("x is greater than y")
    } else if x < y {
        fmt.Println("x is less than y")
    } else {
        fmt.Println("x is equal to y")
    }

    // 5. 循环
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // 6. 函数
    result := add(x, y)
    fmt.Println("Result of add function =", result)
}

func add(a, b int) int {
    return a + b
}