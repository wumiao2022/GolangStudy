package main

import "fmt"

func main() {
    // 实例1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 实例2：计算两个整数的和
    num1 := 10
    num2 := 20
    sum := num1 + num2
    fmt.Println("Sum:", sum)

    // 实例3：判断一个数是奇数还是偶数
    num := 7
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 实例4：计算斐波那契数列
    count := 10
    var fib []int
    fib = append(fib, 0, 1)
    for i := 2; i < count; i++ {
        fib = append(fib, fib[i-1]+fib[i-2])
    }
    fmt.Println("Fibonacci Sequence:", fib)
}