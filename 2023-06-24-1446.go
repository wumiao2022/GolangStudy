package main

import "fmt"

func main() {
    // 练习1: 输出“Hello, World!”
    fmt.Println("Hello, World!")

    // 练习2: 计算并输出2+3的结果
    fmt.Println(2+3)

    // 练习3: 将字符串 "Go is awesome!" 赋值给变量message，并输出该变量的值
    message := "Go is awesome!"
    fmt.Println(message)

    // 练习4: 声明一个整型变量x并赋值为5，然后使用 if 判断 x 是否大于 10。如果大于则输出"greater than 10"，否则输出"less than 10"
    x := 5
    if x > 10 {
        fmt.Println("greater than 10")
    } else {
        fmt.Println("less than 10")
    }

    // 练习5: 使用 for 循环输出 0~9 之间的所有整数
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }
}