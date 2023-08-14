package main

import "fmt"

func main() {
    // 练习1: 打印1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习2: 打印出10的阶乘
    result := 1
    for i := 1; i <= 10; i++ {
        result *= i
    }
    fmt.Println(result)

    // 练习3: 判断一个数是否为素数
    num := 13
    isPrime := true
    for i := 2; i < num; i++ {
        if num%i == 0 {
            isPrime = false
            break
        }
    }
    fmt.Println(isPrime)

    // 练习4: 打印Fibonacci数列的前10个数
    n := 10
    a, b := 0, 1
    for i := 0; i < n; i++ {
        fmt.Println(a)
        a, b = b, a+b
    }
}