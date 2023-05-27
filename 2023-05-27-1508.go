package main

import "fmt"

func main() {
    // 练习1：输出Hello world
    fmt.Println("Hello world")

    // 练习2：实现一个简单的计算器，支持加减乘除四则运算
    var num1, num2 float64
    var operator string
    fmt.Print("请输入第一个数：")
    fmt.Scanln(&num1)
    fmt.Print("请输入第二个数：")
    fmt.Scanln(&num2)
    fmt.Print("请输入操作符(+ - * /)：")
    fmt.Scanln(&operator)
    switch operator {
    case "+":
        fmt.Printf("%.2f + %.2f = %.2f", num1, num2, num1+num2)
    case "-":
        fmt.Printf("%.2f - %.2f = %.2f", num1, num2, num1-num2)
    case "*":
        fmt.Printf("%.2f * %.2f = %.2f", num1, num2, num1*num2)
    case "/":
        if num2 == 0 {
            fmt.Println("除数不能为零")
        } else {
            fmt.Printf("%.2f / %.2f = %.2f", num1, num2, num1/num2)
        }
    default:
        fmt.Println("操作符无效")
    }

    // 练习3：实现一个冒泡排序算法
    arr := []int{3, 9, 4, 7, 2, 8, 1, 5, 6}
    for i := 0; i < len(arr)-1; i++ {
        for j := 0; j < len(arr)-1-i; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    fmt.Println(arr)
}