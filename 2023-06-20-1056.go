package main

import "fmt"

func main() {
    // 1. 变量定义和赋值
    var a int = 10
    b := "hello world"
    c, d := true, false

    // 2. 数组的定义和遍历
    arr := [5]int{1, 2, 3, 4, 5}
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }

    // 3. 切片的定义和遍历
    slice := []int{1, 2, 3, 4, 5}
    for _, v := range slice {
        fmt.Println(v)
    }

    // 4. map的定义和遍历
    m := make(map[string]int)
    m["a"] = 1
    m["b"] = 2
    m["c"] = 3
    for k, v := range m {
        fmt.Println(k, v)
    }

    // 5. if语句
    if a > 0 {
        fmt.Println("a is greater than 0")
    }

    // 6. for循环
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)

    // 7. switch语句
    switch a {
    case 1:
        fmt.Println("a is equal to 1")
    case 2:
        fmt.Println("a is equal to 2")
    default:
        fmt.Println("a is not equal to 1 or 2")
    }

    // 8. 函数定义和调用
    fmt.Println(add(1, 2))

    // 9. defer语句
    defer fmt.Println("world")
    fmt.Println("hello")

    // 10. Panic和Recover
    defer func() {
        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }()
    panic("something went wrong")
}

func add(a, b int) int {
    return a + b
}