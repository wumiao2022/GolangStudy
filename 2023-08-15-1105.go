package main

import "fmt"

func main() {
    // 练习1: 打印出1到10的整数
    for i := 1; i <= 10; i++ {
        fmt.Printf("%d ", i)
    }
    fmt.Println()

    // 练习2: 打印出10到1的整数
    for i := 10; i >= 1; i-- {
        fmt.Printf("%d ", i)
    }
    fmt.Println()

    // 练习3: 打印出0到100之间的偶数
    for i := 0; i <= 100; i += 2 {
        fmt.Printf("%d ", i)
    }
    fmt.Println()

    // 练习4: 打印出100到0之间的奇数
    for i := 99; i >= 1; i -= 2 {
        fmt.Printf("%d ", i)
    }
    fmt.Println()

    // 练习5: 计算并打印1到100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println(sum)
}