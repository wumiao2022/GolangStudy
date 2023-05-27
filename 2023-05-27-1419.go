package main

import (
    "fmt"
)

func main() {
    // 示例1：打印九九乘法表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d*%d=%d\t", j, i, i*j)
        }
        fmt.Println()
    }

    // 示例2：计算斐波那契数列的前20项
    a, b := 1, 1
    fmt.Printf("%d %d ", a, b)
    for i := 3; i <= 20; i++ {
        a, b = b, a+b
        fmt.Printf("%d ", b)
    }

    // 示例3：判断一个字符串是否是回文字符串
    s := "上海自来水来自海上"
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-i-1] {
            fmt.Println("不是回文字符串")
            return
        }
    }
    fmt.Println("是回文字符串")
}