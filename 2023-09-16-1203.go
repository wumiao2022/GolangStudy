package main

import "fmt"

func main() {
    // 题目：计算两个整数的和
    var a, b int = 10, 20
    sum := a + b
    fmt.Println("Sum:", sum)

    // 题目：判断一个数字是否为偶数
    num := 25
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 题目：打印1到10之间的所有数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 题目：判断一个年份是否为闰年
    year := 2024
    if year%4 == 0 {
        if year%100 == 0 {
            if year%400 == 0 {
                fmt.Println(year, "is a leap year")
            } else {
                fmt.Println(year, "is not a leap year")
            }
        } else {
            fmt.Println(year, "is a leap year")
        }
    } else {
        fmt.Println(year, "is not a leap year")
    }
}
