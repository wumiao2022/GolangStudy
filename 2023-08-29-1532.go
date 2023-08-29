package main

import "fmt"

func main() {
    // 打印"Hello, world!"
    fmt.Println("Hello, world!")

    // 定义一个整数变量x，并初始化为10
    x := 10

    // 定义一个字符串变量name，并初始化为"Gopher"
    name := "Gopher"

    // 打印x和name的值
    fmt.Println(x)
    fmt.Println(name)

    // 定义一个切片slice，包含三个整数元素
    slice := []int{1, 2, 3}

    // 打印slice的长度和容量
    fmt.Println(len(slice))
    fmt.Println(cap(slice))

    // 使用for循环遍历slice，并打印每个元素的值
    for _, value := range slice {
        fmt.Println(value)
    }
}