package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // 生成一个随机数种子
    rand.Seed(time.Now().UnixNano())

    // 生成一个 1-100 之间的随机整数
    targetNum := rand.Intn(100) + 1

    guess := 0
    numGuesses := 0

    fmt.Println("我正在想一个 1-100 之间的整数，你可以猜测这个数字")

    // 循环猜数字
    for guess != targetNum {
        fmt.Print("请输入你的猜测：")
        fmt.Scan(&guess)

        if guess < targetNum {
            fmt.Println("你猜的数字太小了，再试试大一点的数字")
        } else if guess > targetNum {
            fmt.Println("你猜的数字太大了，再试试小一点的数字")
        } else {
            fmt.Println("你猜对了！")
            fmt.Printf("你一共猜了 %d 次\n", numGuesses+1)
        }

        numGuesses++
    }
}