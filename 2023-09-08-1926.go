package main

import "fmt"

func main() {
    // 练习1: 打印Hello World
    fmt.Println("Hello, World!")

    // 练习2: 计算两个整数的和并输出
    a := 5
    b := 3
    sum := a + b
    fmt.Println("Sum:", sum)

    // 练习3: 判断一个数是否为偶数
    num := 16
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 练习4: 创建一个字符串切片，并遍历输出其中的元素
    fruits := []string{"apple", "banana", "orange"}
    for _, fruit := range fruits {
        fmt.Println(fruit)
    }

    // 练习5: 使用循环输出1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}