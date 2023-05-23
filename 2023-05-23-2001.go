package main

import (
    "fmt"
)

func main() {
    // 1. 输出Hello, World!
    fmt.Println("Hello, World!")

    // 2. 变量定义和赋值
    var a int = 10
    fmt.Println(a)
    var b = "Golang"
    fmt.Println(b)
    c := true
    fmt.Println(c)

    // 3. 数组基本操作
    arr := []int{1, 2, 3, 4, 5}
    fmt.Println(arr)
    fmt.Println(arr[2])
    fmt.Println(len(arr))

    // 4. 切片基本操作
    slice := arr[1:3]
    fmt.Println(slice)
    slice = append(slice, 6)
    fmt.Println(slice)

    // 5. 控制流语句
    if a > 0 {
        fmt.Println("a is positive")
    } else {
        fmt.Println("a is non-positive")
    }
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }
    i := 0
    for i < len(arr) {
        fmt.Println(arr[i])
        i++
    }
}