package main

import "fmt"

func main() {
    // 1. Hello World
    fmt.Println("Hello, World!")

    // 2. 变量声明与赋值
    var a int
    a = 10
    var b float64 = 3.14
    c := "Hello"
    fmt.Println(a, b, c)

    // 3. 数组遍历
    arr := []int{1, 2, 3, 4, 5}
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }

    // 4. 切片操作
    arr2 := []int{1, 2, 3, 4, 5}
    slc := arr2[1:3]
    fmt.Println(slc)

    // 5. Map遍历
    m := map[string]string{"name": "Tom", "age": "18"}
    for key, value := range m {
        fmt.Println(key, value)
    }

    // 6. if-else条件语句
    score := 90
    if score >= 90 {
        fmt.Println("优秀")
    } else if score >= 60 {
        fmt.Println("及格")
    } else {
        fmt.Println("不及格")
    }

    // 7. for循环
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // 8. switch语句
    num := 1
    switch num {
    case 1:
        fmt.Println("一")
    case 2:
        fmt.Println("二")
    default:
        fmt.Println("其他")
    }

    // 9. 函数的定义和调用
    res := add(2, 3)
    fmt.Println(res)
}

func add(a, b int) int {
    return a + b
}