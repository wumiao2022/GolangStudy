package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    // 生成一个 1 到 100 之间的随机正整数
    randomNumber := rand.Intn(100) + 1

    fmt.Printf("随机生成的数字是：%d\n", randomNumber)
}