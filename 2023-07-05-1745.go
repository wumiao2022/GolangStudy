package main

import "fmt"

func main() {
    // 练习1：打印出Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：定义一个整数变量x，赋值为10，并打印出x的值
    x := 10
    fmt.Println(x)

    // 练习3：计算并打印出1+2+...+10的结果
    sum := 0
    for i := 1; i <= 10; i++ {
        sum += i
    }
    fmt.Println(sum)

    // 练习4：定义一个字符串数组，包含三个姓名，并打印出每个姓名
    names := [3]string{"Alice", "Bob", "Charlie"}
    for _, name := range names {
        fmt.Println(name)
    }

    // 练习5：定义一个函数add，接收两个整数参数，并返回它们的和
    fmt.Println(add(2, 3))
}

func add(a, b int) int {
    return a + b
}