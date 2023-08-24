package main

import (
    "fmt"
)

func main() {
    // 练习1：使用for循环打印数字1到10
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习2：使用for循环计算1到100之间的所有整数的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println(sum)

    // 练习3：使用for循环判断一个数是否为素数
    num := 17
    isPrime := true
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            isPrime = false
            break
        }
    }
    fmt.Println(isPrime)
}