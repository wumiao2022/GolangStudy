package main

import (
    "fmt"
)

func main() {
    // 1. 打印出1-100之间所有的奇数
    for i := 1; i <= 100; i++ {
        if i%2 == 1 {
            fmt.Println(i)
        }
    }

    // 2. 求出斐波那契数列的前20项
    var a, b int = 0, 1
    for i := 0; i < 20; i++ {
        fmt.Println(a)
        a, b = b, a+b
    }

    // 3. 判断一个数是否是素数
    var num int
    fmt.Println("输入一个正整数： ")
    fmt.Scanln(&num)
    if num <= 1 {
        fmt.Printf("%d不是素数", num)
        return
    }
    for i := 2; i <= num/2; i++ {
        if num%i == 0 {
            fmt.Printf("%d不是素数", num)
            return
        }
    }
    fmt.Printf("%d是素数", num)
}