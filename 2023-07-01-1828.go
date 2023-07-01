package main

import "fmt"

func main() {
    // 1. 输出Hello, World!
    fmt.Println("Hello, World!")

    // 2. 计算1+2的结果并输出
    sum := 1 + 2
    fmt.Println(sum)

    // 3. 定义一个整数变量x，并输出其值
    x := 10
    fmt.Println(x)

    // 4. 使用循环打印1到10的所有整数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 5. 定义一个字符串变量message，并初始化为"Hello"
    message := "Hello"
    fmt.Println(message)

    // 6. 定义一个切片，包含整数1到5，并输出
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println(numbers)

    // 7. 使用条件语句判断一个整数x是否为偶数，并输出结果
    if x%2 == 0 {
        fmt.Println("x is even")
    } else {
        fmt.Println("x is odd")
    }

    // 8. 定义一个函数，接收两个整数参数并返回它们的和
    add := func(a, b int) int {
        return a + b
    }
    
    result := add(3, 5)
    fmt.Println(result)
}
