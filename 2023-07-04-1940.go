package main

import (
    "fmt"
)

func main() {
    // 1. 打印 "Hello, World!"
    fmt.Println("Hello, World!")

    // 2. 定义一个整型变量x，赋值为10，并打印出来
    var x int = 10
    fmt.Println(x)

    // 3. 定义一个字符串变量name，赋值为"John"，并打印出来
    name := "John"
    fmt.Println(name)

    // 4. 定义一个切片slice，包含整型元素1, 2, 3，并打印出来
    slice := []int{1, 2, 3}
    fmt.Println(slice)

    // 5. 使用for循环打印出1到10的所有整数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}