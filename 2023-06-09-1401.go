package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano()) // 设置随机数种子
    target := rand.Intn(100)         // 生成随机数

    fmt.Println("猜数字游戏开始！")
    fmt.Println("请输入一个0到100之间的整数：")

    var guess int
    for i := 1; i <= 10; i++ {
        fmt.Printf("第%d次猜测：", i)
        fmt.Scan(&guess)
        if guess == target {
            fmt.Println("恭喜你猜对了！")
            return
        } else if guess < target {
            fmt.Println("猜小了，请继续猜测：")
        } else {
            fmt.Println("猜大了，请继续猜测：")
        }
    }

    fmt.Println("很遗憾，你没有在规定次数内猜中。答案是", target)
}