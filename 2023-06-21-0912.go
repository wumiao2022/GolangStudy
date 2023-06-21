package main

import (
    "fmt"
    "math"
)

func main() {
    // 计算圆的面积和周长
    var r float64 = 5
    var area float64 = math.Pi * r * r
    var circumference float64 = 2 * math.Pi * r
    fmt.Printf("半径为 %.2f 的圆的面积为 %.2f，周长为 %.2f。\n", r, area, circumference)

    // 计算斐波那契数列的前 10 项
    var n int = 10
    var a, b int = 0, 1
    for i := 0; i < n; i++ {
        fmt.Printf("%d ", a)
        a, b = b, a+b
    }

    // 使用字符串切片实现栈结构
    var stack []string
    stack = append(stack, "a")
    stack = append(stack, "b")
    stack = append(stack, "c")
    fmt.Println(stack[len(stack)-1])
    stack = stack[:len(stack)-1]
    fmt.Println(stack[len(stack)-1])
}