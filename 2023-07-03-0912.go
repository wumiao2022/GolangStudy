package main

import "fmt"

func main() {
    // 练习1：将两个整数相加并打印结果
    num1 := 5
    num2 := 3
    sum := num1 + num2
    fmt.Println(sum)

    // 练习2：将两个浮点数相乘并打印结果
    num3 := 3.14
    num4 := 2.5
    product := num3 * num4
    fmt.Println(product)

    // 练习3：判断一个数是否是偶数，并打印结果
    num5 := 10
    if num5%2 == 0 {
        fmt.Println("偶数")
    } else {
        fmt.Println("奇数")
    }

    // 练习4：计算斐波那契数列的第n项，并打印结果
    n := 8
    fib := fibonacci(n)
    fmt.Println(fib)
}

// 斐波那契数列递归函数
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}