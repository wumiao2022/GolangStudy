package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // 生成一个随机数
    rand.Seed(time.Now().UnixNano())
    target := rand.Intn(100)

    // 用户循环猜测直到猜中
    fmt.Println("猜测一个0到99之间的整数")
    fmt.Print("你的猜测：")
    var guess int
    for {
        fmt.Scan(&guess)
        if guess == target {
            fmt.Println("恭喜你猜中了！")
            break
        } else if guess < target {
            fmt.Print("猜小了，再猜：")
        } else {
            fmt.Print("猜大了，再猜：")
        }
    }
}