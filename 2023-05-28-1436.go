package main

import "fmt"

func main() {
    // 实现一个简单的计算器，支持加减乘除四则运算

    num1 := 10
    num2 := 3

    // 加法
    result := num1 + num2
    fmt.Println("加法：", result)

    // 减法
    result = num1 - num2
    fmt.Println("减法：", result)

    // 乘法
    result = num1 * num2
    fmt.Println("乘法：", result)

    // 除法
    result = num1 / num2
    fmt.Println("除法：", result)

    // 取模运算
    result = num1 % num2
    fmt.Println("取模运算：", result)

}