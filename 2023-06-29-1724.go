package main

import "fmt"

func main() {
    // 练习1：打印数字1到10
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习2：打印1到100之间的偶数
    for i := 2; i <= 100; i += 2 {
        fmt.Println(i)
    }

    // 练习3：计算并打印1到100之间的奇数之和
    sum := 0
    for i := 1; i <= 100; i += 2 {
        sum += i
    }
    fmt.Println("Sum:", sum)

    // 练习4：打印一个3x3的九宫格
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            fmt.Print("X ")
        }
        fmt.Println()
    }

    // 练习5：使用嵌套循环打印一个倒立的等腰三角形
    for i := 5; i >= 1; i-- {
        for j := 1; j <= i; j++ {
            fmt.Print("* ")
        }
        fmt.Println()
    }
}