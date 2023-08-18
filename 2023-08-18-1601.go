package main

import "fmt"

func main() {
    // 练习1：打印 Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：打印 1 到 10 的平方
    for i := 1; i <= 10; i++ {
        fmt.Println(i * i)
    }

    // 练习3：计算两个数之和
    a := 5
    b := 10
    sum := a + b
    fmt.Println("The sum of", a, "and", b, "is", sum)
}