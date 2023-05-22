package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    //生成一个随机整数
    num := rand.Intn(100)

    //猜数游戏
    var guess int
    fmt.Println("猜一个0-99之间的整数")
    for {
        fmt.Print("请输入你的猜测：")
        fmt.Scan(&guess)

        if guess < num {
            fmt.Println("猜小了！")
        } else if guess > num {
            fmt.Println("猜大了！")
        } else {
            fmt.Println("猜对了！")
            break
        }
    }
}