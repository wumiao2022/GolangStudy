package main

import "fmt"

func main() {
    // 练习1：打印1到10的所有整数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习2：打印100到1的所有偶数
    for i := 100; i >= 1; i-- {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }

    // 练习3：计算1到100的所有整数的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println(sum)

    // 练习4：计算1到100的所有偶数的和
    sum = 0
    for i := 1; i <= 100; i++ {
        if i%2 == 0 {
            sum += i
        }
    }
    fmt.Println(sum)
}