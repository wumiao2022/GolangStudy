package main

import (
    "fmt"
)

func main() {
    // 练习1: 打印数字1-10
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习2: 打印字符串"Hello, Golang!"
    fmt.Println("Hello, Golang!")

    // 练习3: 计算2的10次方
    result := 1
    for i := 0; i < 10; i++ {
        result *= 2
    }
    fmt.Println(result)

    // 练习4: 判断一个数字是奇数还是偶数
    num := 7
    if num%2 == 0 {
        fmt.Println("偶数")
    } else {
        fmt.Println("奇数")
    }
}
