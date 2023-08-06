package main

import "fmt"

func main() {
    // 练习1：打印出1到10的整数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习2：计算1到100之间所有整数的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("Sum:", sum)

    // 练习3：打印出1到100之间所有奇数
    for i := 1; i <= 100; i += 2 {
        fmt.Println(i)
    }

    // 练习4：输出九九乘法表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d*%d=%d\t", i, j, i*j)
        }
        fmt.Println()
    }

    // 练习5：反转字符串
    str := "Hello, World!"
    reversed := ""
    for _, char := range str {
        reversed = string(char) + reversed
    }
    fmt.Println("Reversed:", reversed)
}
