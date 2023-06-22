package main

import (
    "fmt"
)

func main() {
    // 练习1：打印"Hello, Golang!"
    fmt.Println("Hello, Golang!")

    // 练习2：计算1 + 2 + 3 + ... + 100的结果
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("1 + 2 + 3 + ... + 100 =", sum)

    // 练习3：打印出10以内的所有偶数
    for i := 0; i <= 10; i+=2 {
        fmt.Println(i)
    }

    // 练习4：打印出100以内能够被3整除的数
    for i := 1; i <= 100; i++ {
        if i%3 == 0 {
            fmt.Println(i)
        }
    }
}