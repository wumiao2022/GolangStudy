package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：计算两个整数的和
    num1 := 10
    num2 := 20
    sum := num1 + num2
    fmt.Println("Sum:", sum)

    // 练习3：计算一个整数的平方
    number := 5
    square := number * number
    fmt.Println("Square:", square)

    // 练习4：使用循环打印数字1到10
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习5：判断一个数是否为偶数
    num := 6
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }
}