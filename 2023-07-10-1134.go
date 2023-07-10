package main

import (
    "fmt"
)

func main() {
    // 练习1: 打印出Hello, World!
    fmt.Println("Hello, World!")

    // 练习2: 将两个整数相加并打印出结果
    num1 := 10
    num2 := 5
    sum := num1 + num2
    fmt.Println("Sum:", sum)

    // 练习3: 输出1到10的所有整数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习4: 创建一个切片，并在终端上打印出切片的长度和容量
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println("Length:", len(slice))
    fmt.Println("Capacity:", cap(slice))

    // 练习5: 定义一个函数，将两个整数相乘并返回结果
    multiplyResult := multiply(3, 4)
    fmt.Println("Result:", multiplyResult)
}

func multiply(a, b int) int {
    return a * b
}