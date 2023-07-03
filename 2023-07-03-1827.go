package main

import (
    "fmt"
)

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")
    
    // 练习2：计算1+2+...+10
    sum := 0
    for i := 1; i <= 10; i++ {
        sum += i
    }
    fmt.Println("Sum:", sum)
    
    // 练习3：判断一个数字是否为奇数
    num := 7
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }
    
    // 练习4：反转一个字符串
    str := "Hello, World!"
    reversed := ""
    for _, char := range str {
        reversed = string(char) + reversed
    }
    fmt.Println("Reversed:", reversed)
}