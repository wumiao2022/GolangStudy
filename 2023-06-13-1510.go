package main

import (
    "fmt"
)

func main() {
    // 练习1：输出 "Hello, World!"
    fmt.Println("Hello, World!")

    // 练习2：输入两个整数，输出它们的和
    var a, b int
    fmt.Scan(&a, &b)
    fmt.Println(a + b)

    // 练习3：判断一个数是否为偶数，是则输出 "even"，否则输出 "odd"
    var num int
    fmt.Scan(&num)
    if num % 2 == 0 {
        fmt.Println("even")
    } else {
        fmt.Println("odd")
    }

    // 练习4：定义一个包含10个元素的整数数组，将数组中每个元素乘以2，然后输出数组
    arr := [10]int{1,2,3,4,5,6,7,8,9,10}
    for i := 0; i < len(arr); i++ {
        arr[i] *= 2
        fmt.Print(arr[i], " ")
    }

    // 练习5：定义一个函数，接收一个整数参数，返回该整数的阶乘
    fmt.Println(factorial(5))
}

func factorial(n int) int {
    if n == 1 {
        return 1
    }
    return n * factorial(n-1)
}