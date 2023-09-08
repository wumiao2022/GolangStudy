package main

import "fmt"

func main() {
    // 练习1: 输出 "Hello, World!"
    fmt.Println("Hello, World!")

    // 练习2: 定义一个变量x，并将其初始化为10，然后输出x的值
    var x int = 10
    fmt.Println(x)

    // 练习3: 定义一个切片s，其中包含元素1, 2, 3, 4, 5，并输出切片的长度和容量
    s := []int{1, 2, 3, 4, 5}
    fmt.Println("Length:", len(s))
    fmt.Println("Capacity:", cap(s))

    // 练习4: 定义一个函数add，该函数接受两个整数作为参数，并返回它们的和
    fmt.Println(add(3, 4))

    // 练习5: 使用循环打印出1到10的所有偶数
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }
}

func add(a, b int) int {
    return a + b
}