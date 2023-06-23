package main

import "fmt"

func main() {
    // 打印1到100的整数
    for i := 1; i <= 100; i++ {
        fmt.Println(i)
    }

    // 打印九九乘法口诀表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d*%d=%d\t", j, i, i*j)
        }
        fmt.Println()
    }

    // 打印斐波那契数列前50项
    a, b := 1, 1
    for i := 1; i <= 50; i++ {
        fmt.Println(a)
        a, b = b, a+b
    }

    // 求100内的质数
    for i := 2; i <= 100; i++ {
        isPrime := true
        for j := 2; j < i; j++ {
            if i%j == 0 {
                isPrime = false
                break
            }
        }
        if isPrime {
            fmt.Println(i)
        }
    }
}