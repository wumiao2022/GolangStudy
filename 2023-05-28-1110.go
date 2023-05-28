package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano()) // 初始化随机数种子
    for i := 1; i <= 10; i++ {
        n := rand.Intn(100)
        fmt.Printf("%d ", n)
    }
    fmt.Println()
}

// 生成10个不大于100的随机数打印出来，每个数之间有一个空格。