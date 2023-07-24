package main

import (
    "fmt"
)

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：声明一个变量x并赋值为10，打印变量x的值
    x := 10
    fmt.Println(x)

    // 练习3：声明一个切片s并初始化为[1, 2, 3, 4, 5]，使用循环打印切片中的每个元素
    s := []int{1, 2, 3, 4, 5}
    for i := 0; i < len(s); i++ {
        fmt.Println(s[i])
    }

    // 练习4：声明一个map m并初始化为{"name": "John", "age": 25}，打印map中的键值对
    m := map[string]interface{}{
        "name": "John",
        "age": 25,
    }
    for key, value := range m {
        fmt.Printf("Key: %s, Value: %v\n", key, value)
    }

    // 练习5：实现一个简单的计算器函数，可以进行加、减、乘、除运算
    result := calculator(10, 5, "+")
    fmt.Println(result) // 输出 15

    result = calculator(10, 5, "-")
    fmt.Println(result) // 输出 5

    result = calculator(10, 5, "*")
    fmt.Println(result) // 输出 50

    result = calculator(10, 5, "/")
    fmt.Println(result) // 输出 2
}

func calculator(a, b int, operator string) int {
    switch operator {
    case "+":
        return a + b
    case "-":
        return a - b
    case "*":
        return a * b
    case "/":
        return a / b
    default:
        return 0
    }
}