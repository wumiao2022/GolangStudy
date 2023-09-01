package main

import (
    "fmt"
)

func main() {
    // 练习1：打印1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习2：计算1到100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println(sum)

    // 练习3：判断一个数是否为素数
    num := 17
    isPrime := true
    for i := 2; i < num; i++ {
        if num%i == 0 {
            isPrime = false
            break
        }
    }
    fmt.Println(isPrime)
}