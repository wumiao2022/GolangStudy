package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：打印九九乘法表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d * %d = %d\t", j, i, j*i)
        }
        fmt.Println()
    }

    // 练习3：判断一个数是否为质数
    num := 7
    isPrime := true
    for i := 2; i <= num/2; i++ {
        if num%i == 0 {
            isPrime = false
            break
        }
    }
    if isPrime {
        fmt.Println(num, "is a prime number")
    } else {
        fmt.Println(num, "is not a prime number")
    }
}
