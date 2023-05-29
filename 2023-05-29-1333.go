package main

import (
    "fmt"
    "time"
)

func main() {
    // 实例1：计算1~1000的和
    sum := 0
    for i := 1; i <= 1000; i++ {
        sum += i
    }
    fmt.Println("1~1000的和为：", sum)

    // 实例2：输出当前时间
    now := time.Now()
    fmt.Println("当前时间为：", now.Format("2006-01-02 15:04:05"))

    // 实例3：判断一个数是否为素数
    num := 37
    isPrime := true
    for i := 2; i <= num/2; i++ {
        if num%i == 0 {
            isPrime = false
            break
        }
    }
    if isPrime {
        fmt.Println(num, "是素数")
    } else {
        fmt.Println(num, "不是素数")
    }
}