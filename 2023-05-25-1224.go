package main

import (
    "fmt"
)

func main() {
    // 练习1：输出Hello, Golang!
    fmt.Println("Hello, Golang!")

    // 练习2：计算5+3的结果并输出
    result := 5 + 3
    fmt.Println(result)

    // 练习3：声明一个变量name并赋值为"张三"，输出"我的名字是张三"
    name := "张三"
    fmt.Printf("我的名字是%s\n", name)

    // 练习4：声明一个整型数组[5]int并赋值为[1, 2, 3, 4, 5]，遍历数组并输出每个元素
    arr := [5]int{1, 2, 3, 4, 5}
    for _, v := range arr {
        fmt.Println(v)
    }

    // 练习5：声明一个map[string]string并赋值为{"name": "李四", "age": "18"}，输出map中的每个键值对
    m := map[string]string{"name": "李四", "age": "18"}
    for k, v := range m {
        fmt.Printf("%s: %s\n", k, v)
    }
}