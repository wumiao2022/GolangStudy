package main

import "fmt"

func main() {
    // 1. 变量声明与并行赋值
    var a, b int
    a, b = 1, 2
    fmt.Println(a, b)

    // 2. 条件语句
    if a < b {
        fmt.Println("a is less than b")
    } else if a > b {
        fmt.Println("a is greater than b")
    } else {
        fmt.Println("a equals b")
    }

    // 3. 循环语句
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // 4. 数组
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println(arr)

    // 5. 切片
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println(slice)

    // 6. Map
    m := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }
    fmt.Println(m)

    // 7. 函数
    sum := add(1, 2)
    fmt.Println(sum)
}

func add(a, b int) int {
    return a + b
}