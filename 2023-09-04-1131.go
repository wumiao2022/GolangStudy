package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：反转字符串
    str := "Hello, World!"
    reverseStr := ""
    for i := len(str)-1; i >= 0; i-- {
        reverseStr += string(str[i])
    }
    fmt.Println(reverseStr)

    // 练习3：计算两个数的加法
    num1 := 10
    num2 := 20
    sum := num1 + num2
    fmt.Println(sum)

    // 练习4：打印1到10的奇数
    for i := 1; i <= 10; i++ {
        if i % 2 == 1 {
            fmt.Println(i)
        }
    }
}