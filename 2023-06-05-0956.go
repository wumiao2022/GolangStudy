package main

import "fmt"

func main() {
    // 打印字符串
    fmt.Println("Hello, World!")
    
    // 变量和常量定义
    var x int = 100
    const y string = "Hello"
    
    // if语句
    if x > 0 {
        fmt.Println("x is positive")
    } else if x == 0 {
        fmt.Println("x is zero")
    } else {
        fmt.Println("x is negative")
    }
    
    // for循环
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
    
    // 数组定义和赋值
    var arr [5]int
    arr[0] = 1
    arr[1] = 2
    arr[2] = 3
    arr[3] = 4
    arr[4] = 5
    
    // 遍历数组
    for i := 0; i < len(arr); i++ {
        fmt.Printf("arr[%d] = %d\n", i, arr[i])
    }
    
    // 切片定义和赋值
    var slice []int = []int{1, 2, 3, 4, 5}
    
    // 遍历切片
    for i, v := range slice {
        fmt.Printf("slice[%d] = %d\n", i, v)
    }
    
    // map定义和赋值
    var m map[string]int = map[string]int{
        "apple":  1,
        "banana": 2,
        "orange": 3,
    }
    
    // 遍历map
    for k, v := range m {
        fmt.Printf("%s -> %d\n", k, v)
    }
    
    // 函数定义和调用
    sum := add(3, 4)
    fmt.Println(sum)
}

// 简单的加法函数
func add(x, y int) int {
    return x + y
}