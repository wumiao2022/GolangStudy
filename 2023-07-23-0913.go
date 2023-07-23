package main

import "fmt"

func main() {
    // 练习1: 打印 Hello, World!
    fmt.Println("Hello, World!")

    // 练习2: 使用 for 循环打印 1 到 10
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习3: 将两个整数相加并打印结果
    a := 5
    b := 3
    sum := a + b
    fmt.Println("Sum:", sum)

    // 练习4: 判断一个数是否为偶数并打印结果
    num := 7
    if num % 2 == 0 {
        fmt.Println("Even")
    } else {
        fmt.Println("Odd")
    }

    // 练习5: 使用切片存储一组整数，并打印出所有元素
    numbers := []int{1, 2, 3, 4, 5}
    for _, num := range numbers {
        fmt.Println(num)
    }
}