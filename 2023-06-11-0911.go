package main

import (
    "fmt"
)

func main() {
    // 计算两个数的和
    num1 := 10
    num2 := 5
    sum := num1 + num2
    fmt.Println("Sum:", sum)

    // 求出1到100的和
    sum2 := 0
    for i := 1; i <= 100; i++ {
        sum2 += i
    }
    fmt.Println("Sum2:", sum2)

    // 判断一个数是否为偶数
    num3 := 7
    if num3%2 == 0 {
        fmt.Println(num3, "is even")
    } else {
        fmt.Println(num3, "is odd")
    }

    // 计算一个数的阶乘
    num4 := 5
    fact := 1
    for i := 1; i <= num4; i++ {
        fact *= i
    }
    fmt.Println(num4, "!= ", fact)

    // 打印一个等边三角形
    rows := 5
    for i := 1; i <= rows; i++ {
        for j := 1; j <= rows-i; j++ {
            fmt.Print(" ")
        }
        for k := 1; k <= i*2-1; k++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}