package main

import "fmt"

func main() {
    // 实现一个简单的计算器
    num1 := 10.0
    num2 := 5.0
    operator := "+"

    switch operator {
    case "+":
        fmt.Println(num1 + num2)
    case "-":
        fmt.Println(num1 - num2)
    case "*":
        fmt.Println(num1 * num2)
    case "/":
        if num2 == 0 {
            fmt.Println("除数不能为零")
        } else {
            fmt.Println(num1 / num2)
        }
    default:
        fmt.Println("无法识别的操作符")
    }

    // 实现一个简单的冒泡排序
    arr := []int{5, 3, 8, 2, 9, 1}

    for i := 0; i < len(arr)-1; i++ {
        for j := 0; j < len(arr)-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    fmt.Println(arr)

    // 实现一个简单的递归函数，求斐波那契数列的第n项
    n := 10
    fmt.Println(fibonacci(n))
}

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}