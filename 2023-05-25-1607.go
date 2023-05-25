package main

import "fmt"

func main() {
    // 练习1：输出"Hello, World!"
    fmt.Println("Hello, World!")

    // 练习2：打印九九乘法表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d*%d=%d\t", i, j, i*j)
        }
        fmt.Println()
    }

    // 练习3：计算斐波那契数列的第n项
    n := 10 // 求第10项
    fmt.Printf("斐波那契数列的第%d项为：%d", n, fibonacci(n))
}

func fibonacci(n int) int {
    if n <= 2 {
        return 1
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
