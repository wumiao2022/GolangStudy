package main

import (
    "fmt"
)

func main() {
    // 1. 输出0到9
    for i := 0; i < 10; i++ {
        fmt.Print(i, " ")
    }
    fmt.Println()

    // 2. 输出1到100中的奇数
    for i := 1; i <= 100; i += 2 {
        fmt.Print(i, " ")
    }
    fmt.Println()

    // 3. 输入一个整数n，输出n的阶乘
    var n, res int = 5, 1
    for i := 1; i <= n; i++ {
        res *= i
    }
    fmt.Println(res)

    // 4. 求Fibonacci数列的前10个数
    var a, b = 0, 1
    for i := 0; i < 10; i++ {
        fmt.Print(a, " ")
        a, b = b, a+b
    }
    fmt.Println()

    // 5. 判断一个数是否为质数
    var num int = 17
    var flag bool = true
    for i := 2; i < num; i++ {
        if num%i == 0 {
            flag = false
            break
        }
    }
    if flag {
        fmt.Println(num, "是质数")
    } else {
        fmt.Println(num, "不是质数")
    }
}