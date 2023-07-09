package main

import (
    "fmt"
    "time"
)

func main() {
    // 练习1：输出Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：输出当前时间
    fmt.Println("Current time:", time.Now())

    // 练习3：计算两个数的和并输出
    a := 5
    b := 10
    sum := a + b
    fmt.Println("Sum:", sum)

    // 练习4：使用循环输出1到10之间的奇数
    for i := 1; i <= 10; i += 2 {
        fmt.Println(i)
    }

    // 练习5：判断一个数是否为偶数并输出结果
    number := 6
    if number%2 == 0 {
        fmt.Println(number, "is even")
    } else {
        fmt.Println(number, "is odd")
    }
}

// 练习6：实现一个简单的计算器，根据用户输入的操作符和两个数进行运算并输出结果
func calculator(operator string, a, b int) int {
    var result int

    switch operator {
    case "+":
        result = a + b
    case "-":
        result = a - b
    case "*":
        result = a * b
    case "/":
        result = a / b
    default:
        fmt.Println("Invalid operator")
    }

    return result
}