package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：计算并打印1+2的结果
    fmt.Println(1 + 2)

    // 练习3：定义一个变量name并赋值为"John"，打印出"Hello, John!"
    name := "John"
    fmt.Println("Hello, " + name + "!")

    // 练习4：定义一个整型数组并打印出数组的长度
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println(len(numbers))

    // 练习5：定义一个函数add，接收两个整型参数并返回它们的和，然后调用并打印结果
    fmt.Println(add(3, 5))
}

func add(a, b int) int {
    return a + b
}
