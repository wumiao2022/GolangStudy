package main

import "fmt"

func main() {
    // 练习1：输出Hello World
    fmt.Println("Hello World")

    // 练习2：变量的声明与赋值
    var num1 int = 10
    var num2 int
    num2 = 20
    fmt.Println("num1:", num1)
    fmt.Println("num2:", num2)

    // 练习3：常量的声明与使用
    const PI float64 = 3.14159
    fmt.Println("PI:", PI)

    // 练习4：判断语句
    score := 80
    if score >= 60 {
        fmt.Println("及格")
    } else {
        fmt.Println("不及格")
    }

    // 练习5：循环语句
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }
}