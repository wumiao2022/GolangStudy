package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // 生成一个随机数，并判断是否为偶数
    rand.Seed(time.Now().UnixNano())
    number := rand.Intn(100)

    if number%2 == 0 {
        fmt.Printf("%d 是偶数\n", number)
    } else {
        fmt.Printf("%d 是奇数\n", number)
    }
}