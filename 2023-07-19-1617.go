package main

import (
    "fmt"
    "math"
)

func main() {
    // 练习1: 计算两个数的和，并输出结果
    num1 := 10
    num2 := 5
    sum := num1 + num2
    fmt.Println("Sum:", sum)
    
    // 练习2: 求1到100的累加和，并输出结果
    total := 0
    for i := 1; i <= 100; i++ {
        total += i
    }
    fmt.Println("Total:", total)

    // 练习3: 计算圆的面积和周长，并输出结果
    radius := 5.0
    area := math.Pi * math.Pow(radius, 2)
    circumference := 2 * math.Pi * radius
    fmt.Printf("Area: %.2f\n", area)
    fmt.Printf("Circumference: %.2f\n", circumference)
    
    // 练习4: 输出螺旋矩阵
    n := 4
    matrix := make([][]int, n)
    for i := 0; i < n; i++ {
        matrix[i] = make([]int, n)
    }
    num := 1
    for layer := 0; layer < n/2+1; layer++ {
        // 上边
        for i := layer; i < n-layer; i++ {
            matrix[layer][i] = num
            num++
        }
        // 右边
        for i := layer + 1; i < n-layer; i++ {
            matrix[i][n-layer-1] = num
            num++
        }
        // 下边
        for i := n - layer - 2; i >= layer; i-- {
            matrix[n-layer-1][i] = num
            num++
        }
        // 左边
        for i := n - layer - 2; i > layer; i-- {
            matrix[i][layer] = num
            num++
        }
    }
    // 输出矩阵
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            fmt.Printf("%d\t", matrix[i][j])
        }
        fmt.Println()
    }
}