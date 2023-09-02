package main

import "fmt"

func main() {
    // 题目：输出 "Hello, World!"
    fmt.Println("Hello, World!")

    // 题目：计算并输出 1+2+3+...+10 的结果
    sum := 0
    for i := 1; i <= 10; i++ {
        sum += i
    }
    fmt.Println("1+2+3+...+10 =", sum)

    // 题目：交换两个变量的值，并输出结果
    a := 5
    b := 10
    a, b = b, a
    fmt.Println("After swapping, a =", a, "and b =", b)

    // 题目：判断一个数是奇数还是偶数，并输出结果
    num := 6
    if num%2 == 0 {
        fmt.Println(num, "is an even number.")
    } else {
        fmt.Println(num, "is an odd number.")
    }

    // 题目：输出 1-10 之间的所有偶数
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }
}