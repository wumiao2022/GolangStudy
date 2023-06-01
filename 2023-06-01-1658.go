package main

import "fmt"

func main() {
    // 练习1：打印出1~100中的奇数
    for i := 1; i <= 100; i++ {
        if i%2 == 1 {
            fmt.Println(i)
        }
    }

    // 练习2：计算1~100的累加和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("1~100的累加和为：", sum)

    // 练习3：判断一个数是否为素数
    number := 17
    isPrime := true
    for i := 2; i <= number/2; i++ {
        if number%i == 0 {
            isPrime = false
            break
        }
    }
    if isPrime {
        fmt.Println(number, "是素数")
    } else {
        fmt.Println(number, "不是素数")
    }
}