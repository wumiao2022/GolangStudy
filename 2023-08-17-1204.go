package main

import "fmt"

func main() {
    // 实例1: 变量交换
    a := 10
    b := 20
    a, b = b, a
    fmt.Println(a, b)

    // 实例2: 判断一个数是否为偶数
    num := 8
    if num % 2 == 0 {
        fmt.Println("偶数")
    } else {
        fmt.Println("奇数")    
    }

    // 实例3: 使用for循环打印乘法表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d*%d=%d ", j, i, j*i)
        }
        fmt.Println()
    }

    // 实例4: 字符串反转
    str := "Hello, World!"
    reversedStr := ""
    for i := len(str) - 1; i >= 0; i-- {
        reversedStr += string(str[i])
    }
    fmt.Println(reversedStr)
}