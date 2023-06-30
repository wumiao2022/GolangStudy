package main

import "fmt"

func main() {
    // 练习1：输出10到1倒序的数字
    for i := 10; i >= 1; i-- {
        fmt.Println(i)
    }

    // 练习2：计算1到100的累加和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("Sum:", sum)

    // 练习3：输出1到100之间的偶数
    for i := 1; i <= 100; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }
}