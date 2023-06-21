package main

import (
    "fmt"
)

func main() {
    // 求1-100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("1-100的和：", sum)

    // 判断一个数是否是质数
    num := 23
    isPrime := true
    if num <= 1 {
        isPrime = false
    }
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            isPrime = false
            break
        }
    }
    if isPrime {
        fmt.Println(num, "是质数")
    } else {
        fmt.Println(num, "不是质数")
    }

    // 打印一个菱形
    for i := 1; i <= 5; i++ {
        for j := 1; j <= 5-i; j++ {
            fmt.Print(" ")
        }
        for k := 1; k <= 2*i-1; k++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
    for i := 4; i >= 1; i-- {
        for j := 1; j <= 5-i; j++ {
            fmt.Print(" ")
        }
        for k := 1; k <= 2*i-1; k++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}