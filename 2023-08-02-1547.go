package main

import "fmt"

func main() {
    // 练习1: 打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2: 声明并初始化一个整型变量，然后打印该变量的值
    var num int = 10
    fmt.Println(num)

    // 练习3: 通过循环语句打印1到10之间所有的偶数
    for i := 1; i <= 10; i++ {
        if i % 2 == 0 {
            fmt.Print(i, " ")
        }
    }
    fmt.Println()

    // 练习4: 实现一个函数Add，接受两个整数作为参数，返回它们的和
    fmt.Println(Add(2, 3))

    // 练习5: 声明一个切片，包含一些整数，并使用range关键字打印切片中的每个元素
    numbers := []int{1, 2, 3, 4, 5}
    for _, num := range numbers {
        fmt.Print(num, " ")
    }
    fmt.Println()

    // 练习6: 实现一个计算阶乘的函数Factorial，接受一个整数作为参数，返回它的阶乘
    fmt.Println(Factorial(5))
}

func Add(a, b int) int {
    return a + b
}

func Factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * Factorial(n-1)
}