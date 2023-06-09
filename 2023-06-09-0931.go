package main

import "fmt"
import "time"

func main() {
    // 1. 输出Hello World
    fmt.Println("Hello World")

    // 2. 输出当前时间
    fmt.Println(time.Now())

    // 3. 定义变量并输出
    name := "Bob"
    fmt.Println("My name is", name)

    // 4. 判断语句示例
    if 1+1 == 2 {
        fmt.Println("1+1 equals 2")
    } else {
        fmt.Println("Something is wrong")
    }

    // 5. 循环语句示例
    for i := 0; i < 5; i++ {
        fmt.Println("Counting", i+1)
    }
}