package main

import "fmt"

func main() {
    // 1. 打印 Hello, World!
    fmt.Println("Hello, World!")

    // 2. 定义一个整数变量，并打印出来
    num := 5
    fmt.Println(num)

    // 3. 定义一个字符串切片，并打印出切片的长度和容量
    strSlice := []string{"apple", "banana", "cherry"}
    fmt.Println(len(strSlice))
    fmt.Println(cap(strSlice))

    // 4. 使用循环语句打印 1 到 10 的所有偶数
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }

    // 5. 使用 if-else 判断一个数是否为正数、负数或零，并打印结果
    num2 := -3
    if num2 > 0 {
        fmt.Println("Positive number")
    } else if num2 < 0 {
        fmt.Println("Negative number")
    } else {
        fmt.Println("Zero")
    }
}