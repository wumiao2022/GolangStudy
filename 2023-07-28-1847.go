package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：打印1到10之间的所有整数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习3：计算两个数字的和
    num1 := 10
    num2 := 20
    sum := num1 + num2
    fmt.Println("Sum:", sum)

    // 练习4：判断一个数字是否为偶数
    number := 15
    if number % 2 == 0 {
        fmt.Println(number, "is even")
    } else {
        fmt.Println(number, "is odd")
    }

    // 练习5：计算字符串的长度
    str := "Hello, World!"
    length := len(str)
    fmt.Println("Length of str:", length)
}