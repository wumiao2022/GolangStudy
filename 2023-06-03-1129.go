package main

import (
    "fmt"
)

func main() {
    // 1. 打印 "Hello, World!" 字符串
    fmt.Println("Hello, World!")

    // 2. 计算 12 + 36 的结果并打印出来
    fmt.Println(12 + 36)

    // 3. 打印从 1 到 10 的所有整数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 4. 定义一个名为 "message" 的字符串变量，并将其初始化为 "Hello, Golang!"
    message := "Hello, Golang!"
    fmt.Println(message)

    // 5. 定义一个名为 "numbers" 的整型数组，并将其初始化为 1, 2, 3, 4, 5
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println(numbers)

    // 6. 将 "Hello, World!" 字符串反转并输出结果
    message = "Hello, World!"
    reversed := ""
    for i := len(message) - 1; i >= 0; i-- {
        reversed += string(message[i])
    }
    fmt.Println(reversed)

    // 7. 定义一个名为 "square" 的函数，接受一个整型参数并返回该整数的平方
    square := func(n int) int {
        return n * n
    }

    // 8. 将 "I am an expert Golang programmer!" 字符串中的所有大写字母转换成小写字母并输出结果
    message = "I am an expert Golang programmer!"
    lowercase := ""
    for _, c := range message {
        if c >= 'A' && c <= 'Z' {
            lowercase += string(c + 32)
        } else {
            lowercase += string(c)
        }
    }
    fmt.Println(lowercase)

    // 9. 将一个名为 "keys" 的字符串数组按照字母顺序进行排序，并输出结果
    keys := []string{"banana", "apple", "grape", "orange"}
    sort.Strings(keys)
    fmt.Println(keys)

    // 10. 使用 goroutine 运行一个名为 "count" 的函数，该函数接受一个整型参数并连续打印该整数减 1 的结果，直到该整数减为 0
    count := func(n int) {
        for i := n; i > 0; i-- {
            fmt.Println(i - 1)
        }
    }
    go count(10)
}