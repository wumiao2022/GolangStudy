package main

import "fmt"

func main() {
    // 练习1：输出Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：计算1到10的和并输出
    sum := 0
    for i := 1; i <= 10; i++ {
        sum += i
    }
    fmt.Println("Sum:", sum)

    // 练习3：判断一个数是奇数还是偶数并输出
    num := 9
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 练习4：打印九九乘法表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d*%d=%d ", j, i, j*i)
        }
        fmt.Println()
    }

    // 练习5：计算斐波那契数列并输出前20项
    fib := []int{0, 1}
    for i := 2; i < 20; i++ {
        fib = append(fib, fib[i-1]+fib[i-2])
    }
    fmt.Println("Fibonacci sequence:", fib)
}