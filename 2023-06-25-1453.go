package main

import "fmt"

func main() {
    // 1. 输出hello world
    fmt.Println("Hello World!")

    // 2. 变量声明和初始化
    var a int = 10
    var b string = "Go"
    fmt.Println(a)
    fmt.Println(b)

    // 3. 常量声明
    const PI float64 = 3.1415926535
    fmt.Println(PI)

    // 4. 条件语句if
    if a > 0 {
        fmt.Println("a is positive")
    } else if a == 0 {
        fmt.Println("a is zero")
    } else {
        fmt.Println("a is negative")
    }

    // 5. for循环
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println("Sum of first 10 integers is ", sum)

    // 6. switch语句
    switch b {
    case "Go":
        fmt.Println("b is Go")
    case "Java":
        fmt.Println("b is Java")
    default:
        fmt.Println("b is neither Go nor Java")
    }

    // 7. 数组声明和遍历
    var arr [5]int
    for i := 0; i < len(arr); i++ {
        arr[i] = i
    }
    fmt.Println(arr)

    // 8. 切片声明和遍历
    slice := []int{1, 2, 3, 4, 5}
    for _, val := range slice {
        fmt.Println(val)
    }

    // 9. 结构体声明和使用
    type Person struct {
        name   string
        age    int
        gender string
    }
    person := Person{name: "John", age: 25, gender: "Male"}
    fmt.Println(person)

    // 10. 函数声明和使用
    add := func(a, b int) int {
        return a + b
    }
    result := add(2, 3)
    fmt.Println(result)
}