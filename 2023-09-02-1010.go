package main

import "fmt"

func main() {
    // 练习1: 输出Hello, World!
    fmt.Println("Hello, World!")

    // 练习2: 输出1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习3: 计算1到100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("Sum:", sum)

    // 练习4: 判断一个数是否为偶数
    num := 15
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 练习5: 打印九九乘法表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print(j, " * ", i, " = ", j*i, "\t")
        }
        fmt.Println()
    }
}