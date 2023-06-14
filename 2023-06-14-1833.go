package main

import "fmt"

func main() {
    // 1. 变量与数据类型
    var a int = 10
    b := "Hello World"
    c := 3.14

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)

    // 2. 数组和切片
    arr1 := [3]int{1, 2, 3}
    arr2 := [...]int{4, 5, 6}
    slice1 := []string{"a", "b", "c"}

    fmt.Println(arr1)
    fmt.Println(arr2)
    fmt.Println(slice1)

    // 3. 控制流程
    for i := 0; i < 5; i++ {
        if i == 3 {
            continue
        }
        fmt.Println(i)
    }

    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }

    j := 0
    for {
        fmt.Println(j)
        j++
        if j == 5 {
            break
        }
    }

    // 4. 函数
    add := func(a, b int) int {
        return a + b
    }

    result := add(1, 2)
    fmt.Println(result)

    // 5. 包的使用
    fmt.Println("Hello")
}