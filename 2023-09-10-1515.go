package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // 设置随机数种子
    rand.Seed(time.Now().UnixNano())

    // 生成一个随机整数
    randomInt := rand.Intn(100)

    // 打印随机整数的值
    fmt.Println("随机整数:", randomInt)

    // 判断随机整数的奇偶性
    if randomInt%2 == 0 {
        fmt.Println("随机整数为偶数")
    } else {
        fmt.Println("随机整数为奇数")
    }
}