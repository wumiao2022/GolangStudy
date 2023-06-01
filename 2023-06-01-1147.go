package main

import "fmt"

func main() {
    // 练习1：打印1到100的数字
    for i := 1; i <= 100; i++ {
        fmt.Println(i)
    }

    // 练习2：打印九九乘法表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d*%d=%d ", j, i, i*j)
        }
        fmt.Println()
    }
    
    // 练习3：求100以内的素数

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