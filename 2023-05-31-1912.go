package main

import "fmt"

func main() {
    // 打印乘法口诀表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d*%d=%d\t", j, i, i*j)
        }
        fmt.Println()
    }

    // 遍历切片并输出元素
    fruits := []string{"apple", "banana", "orange", "grape"}
    for i, fruit := range fruits {
        fmt.Printf("Fruit %d: %s\n", i+1, fruit)
    }

    // 实现斐波那契数列
    n := 10
    fib := []int{0, 1}
    for i := 2; i < n; i++ {
        fib = append(fib, fib[i-1]+fib[i-2])
    }
    fmt.Println(fib)

    // 实现一个简单的计算器
    var operator string
    var num1, num2 float64
    fmt.Println("请输入要计算的数字和运算符，用空格隔开：")
    fmt.Scanf("%f %s %f", &num1, &operator, &num2)
    switch operator {
    case "+":
        fmt.Printf("结果为：%f\n", num1+num2)
    case "-":
        fmt.Printf("结果为：%f\n", num1-num2)
    case "*":
        fmt.Printf("结果为：%f\n", num1*num2)
    case "/":
        if num2 == 0 {
            fmt.Println("除数不能为0")
            return
        }
        fmt.Printf("结果为：%f\n", num1/num2)
    default:
        fmt.Println("不支持的运算符")
    }
}