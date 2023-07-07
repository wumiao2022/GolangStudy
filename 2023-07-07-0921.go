package main

import "fmt"

func main() {
    // 1. 打印 "Hello, World!"
    fmt.Println("Hello, World!")

    // 2. 声明一个整数变量，并打印出其值
    var i int
    fmt.Println(i)

    // 3. 声明两个字符串变量，并打印出它们的拼接结果
    var str1 string = "Hello"
    var str2 string = "World"
    fmt.Println(str1 + " " + str2)

    // 4. 声明一个整数数组，并打印出数组中的元素
    arr := []int{1, 2, 3, 4, 5}
    for _, num := range arr {
        fmt.Println(num)
    }

    // 5. 声明一个Map，并打印出其中的键值对
    m := map[string]int{"A": 1, "B": 2, "C": 3}
    for k, v := range m {
        fmt.Println(k, v)
    }
}