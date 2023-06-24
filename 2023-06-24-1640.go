package main

import (
    "fmt"
)

func main() {
    // 练习题1：输出1~10的整数
    for i := 1; i <= 10; i++ {
        fmt.Printf("%d ", i)
    }
    fmt.Println()

    // 练习题2：输出10~1的整数
    for i := 10; i >= 1; i-- {
        fmt.Printf("%d ", i)
    }
    fmt.Println()

    // 练习题3：输出1~100之间所有奇数和
    sum := 0
    for i := 1; i <= 100; i++ {
        if i%2 == 1 {
            sum += i
        }
    }
    fmt.Println(sum)

    // 练习题4：输出斐波那契数列前20项
    a, b := 0, 1
    for i := 0; i < 20; i++ {
        fmt.Printf("%d ", a)
        a, b = b, a+b
    }
    fmt.Println()

    // 练习题5：输出乘法口诀表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d*%d=%d ", j, i, i*j)
        }
        fmt.Println()
    }
}