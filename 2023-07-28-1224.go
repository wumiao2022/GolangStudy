package main

import "fmt"

func main() {
    // 练习1：打印出1到10的所有整数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习2：使用循环计算1到100之间所有整数的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("Sum:", sum)

    // 练习3：打印出从A-Z的所有字母
    for c := 'A'; c <= 'Z'; c++ {
        fmt.Printf("%c ", c)
    }

    // 练习4：打印出一个10x10的矩形
    for i := 1; i <= 10; i++ {
        for j := 1; j <= 10; j++ {
            fmt.Print("* ")
        }
        fmt.Println()
    }
}

// PS: 每天多练习才能提升编程技能哦！祝你编程愉快！