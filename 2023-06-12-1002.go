package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano()) // 设置随机数种子

    for i := 0; i < 5; i++ {
        num := rand.Intn(100) // 生成0到99之间的随机整数
        fmt.Printf("随机数%d: ", num)
        if num%2 == 0 {
            fmt.Println("偶数")
        } else {
            fmt.Println("奇数")
        }
    }
}