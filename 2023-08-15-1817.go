package main

import "fmt"

func main() {
    // 练习1: 输出Hello, World!
    fmt.Println("Hello, World!")

    // 练习2: 计算1+2的结果并输出
    fmt.Println(1 + 2)

    // 练习3: 创建一个长度为5的整数数组并输出
    var arr [5]int
    fmt.Println(arr)

    // 练习4: 将数组中的元素赋值为1, 2, 3, 4, 5，并输出
    for i := 0; i < 5; i++ {
        arr[i] = i + 1
    }
    fmt.Println(arr)

    // 练习5: 使用for循环输出数组中的每个元素
    for _, num := range arr {
        fmt.Println(num)
    }
}