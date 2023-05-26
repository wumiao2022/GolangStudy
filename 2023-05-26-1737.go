package main

import (
    "fmt"
)

func main() {
    // 1. 打印Hello World
    fmt.Println("Hello World")

    // 2. 定义并使用变量
    var x int = 5
    fmt.Println("x =", x)

    // 3. 数组的遍历
    arr := [5]int{1, 2, 3, 4, 5}
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }

    // 4. 切片的遍历
    slice := []int{1, 2, 3, 4, 5}
    for i := 0; i < len(slice); i++ {
        fmt.Println(slice[i])
    }

    // 5. map的遍历
    numbers := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5}
    for key, value := range numbers {
        fmt.Println(key, ":", value)
    }

    // 6. if语句的使用
    if x > 3 {
        fmt.Println("x is greater than 3")
    } else {
        fmt.Println("x is less than or equal to 3")
    }

    // 7. for循环的使用
    sum := 0
    for i := 0; i < 5; i++ {
        sum += i
    }
    fmt.Println("sum =", sum)

    // 8. switch语句的使用
    num := 3
    switch num {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    default:
        fmt.Println("other")
    }

    // 9. 函数的定义和调用
    result := add(4, 5)
    fmt.Println("result =", result)

    // 10. 结构体的定义和使用
    person := Person{"John", 25}
    fmt.Println(person)
}

type Person struct {
    name string
    age  int
}

func add(a, b int) int {
    return a + b
}