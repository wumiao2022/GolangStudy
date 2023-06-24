package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    for i := 1; i <= 10; i++ {
        randomNum := rand.Intn(100)
        fmt.Printf("第%d个随机数是：%d\n", i, randomNum)
    }
}