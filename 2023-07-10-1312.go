package main

import (
    "fmt"
)

func main() {
    // 练习1: 打印输出 "Hello, World!"
    fmt.Println("Hello, World!")

    // 练习2: 声明一个整数变量并打印输出
    var num int = 10
    fmt.Println(num)

    // 练习3: 声明一个字符串变量并打印输出
    var str string = "Golang"
    fmt.Println(str)

    // 练习4: 声明一个切片并使用循环输出其中的元素
    nums := []int{1, 2, 3, 4, 5}
    for _, n := range nums {
        fmt.Println(n)
    }

    // 练习5: 声明一个map并打印输出其中的键值对
    info := map[string]string{
        "name":  "Alice",
        "age":   "25",
        "city":  "New York",
        "email": "alice@example.com",
    }
    for key, value := range info {
        fmt.Printf("%s: %s\n", key, value)
    }
}