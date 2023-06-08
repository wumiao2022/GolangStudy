package main

import (
    "fmt"
)

func main() {
    // 1. 输出Hello, world!
    fmt.Println("Hello, world!")

    // 2. 输出10以内的偶数
    for i := 0; i <= 10; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }

    // 3. 计算1到100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println(sum)

    // 4. 输出九九乘法表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d*%d=%d ", j, i, i*j)
        }
        fmt.Println()
    }

    // 5. 判断一个数是否为素数
    num := 13
    isPrime := true
    for i := 2; i < num; i++ {
        if num%i == 0 {
            isPrime = false
            break
        }
    }
    if isPrime {
        fmt.Printf("%d is a prime number\n", num)
    } else {
        fmt.Printf("%d is not a prime number\n", num)
    }
}