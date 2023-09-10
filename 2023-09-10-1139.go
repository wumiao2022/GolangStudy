package main

import (
    "fmt"
    "time"
)

func main() {
    // 练习1: 打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2: 打印当前时间
    fmt.Println("当前时间:", time.Now())

    // 练习3: 声明和初始化一个整数变量并打印
    var num int = 10
    fmt.Println("整数变量num:", num)

    // 练习4: 声明和初始化一个字符串变量并打印
    var str string = "Hello, Golang!"
    fmt.Println("字符串变量str:", str)

    // 练习5: 使用循环计算1到10的和并打印结果
    sum := 0
    for i := 1; i <= 10; i++ {
        sum += i
    }
    fmt.Println("1到10的和:", sum)

    // 练习6: 声明一个切片，并对切片进行追加元素和删除元素的操作
    var slice []int // 声明一个整型切片
    slice = append(slice, 1, 2, 3) // 追加元素1, 2, 3
    fmt.Println("切片初始值:", slice)

    slice = append(slice, 4) // 追加元素4
    fmt.Println("追加元素后的切片:", slice)

    slice = append(slice[:2], slice[3:]...) // 删除索引为2的元素
    fmt.Println("删除元素后的切片:", slice)
}