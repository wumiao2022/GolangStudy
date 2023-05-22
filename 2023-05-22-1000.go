package main

import "fmt"

func main() {
    // 1. 输出 Hello World
    fmt.Println("Hello World")

    // 2. 输出 1~100 的整数
    for i := 1; i <= 100; i++ {
        fmt.Println(i)
    }

    // 3. 打印乘法口诀表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d * %d = %d\t", j, i, i*j)
        }
        fmt.Println()
    }

    // 4. 求出 100 以内的所有质数
    for i := 2; i <= 100; i++ {
        isPrime := true
        for j := 2; j < i; j++ {
            if i%j == 0 {
                isPrime = false
            }
        }
        if isPrime {
            fmt.Println(i)
        }
    }

    // 5. 求出斐波那契数列中的前 20 个数
    a, b := 0, 1
    for i := 1; i <= 20; i++ {
        fmt.Println(a)
        a, b = b, a+b
    }
}