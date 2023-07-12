package main

import "fmt"

func main() {
    // 练习1：打印出1到10的所有整数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    fmt.Println()

    // 练习2：打印出2到20之间的所有偶数
    for i := 2; i <= 20; i += 2 {
        fmt.Println(i)
    }

    fmt.Println()

    // 练习3：使用循环将一个字符串反转并打印
    text := "Hello, world!"
    for i := len(text)-1; i >= 0; i-- {
        fmt.Print(string(text[i]))
    }
    fmt.Println()

    // 练习4：使用嵌套循环打印一个九九乘法表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= 9; j++ {
            fmt.Printf("%d x %d = %d\t", i, j, i*j)
        }
        fmt.Println()
    }
}