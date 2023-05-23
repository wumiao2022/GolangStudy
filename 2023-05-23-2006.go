package main

import "fmt"

func main() {
    // 练习1：输出10以内的奇数
    for i := 1; i <= 10; i += 2 {
        fmt.Println(i)
    }

    fmt.Println()

    // 练习2：输出1到100里，只有3的倍数，但不是9的倍数的数
    for i := 1; i <= 100; i++ {
        if i%3 == 0 && i%9 != 0 {
            fmt.Println(i)
        }
    }

    fmt.Println()

    // 练习3：输出100以内的斐波那契数列
    a, b := 0, 1
    for b < 100 {
        fmt.Println(b)
        a, b = b, a+b
    }
}