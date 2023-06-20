package main

import "fmt"

func main() {
    // 练习1：输出Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：输出1到10的所有整数
    for i := 1; i <= 10; i++ {
        fmt.Print(i, " ")
    }
    fmt.Println()

    // 练习3：计算5的阶乘
    result := 1
    for i := 1; i <= 5; i++ {
        result *= i
    }
    fmt.Println("5的阶乘为：", result)

    // 练习4：定义一个数组，包含1到5的所有整数，并输出数组中的所有元素
    arr := [...]int{1, 2, 3, 4, 5}
    for _, v := range arr {
        fmt.Print(v, " ")
    }
    fmt.Println()

    // 练习5：计算两个整数的和
    a, b := 3, 5
    sum := a + b
    fmt.Println("3 + 5 = ", sum)
}