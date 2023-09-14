package main

import (
    "fmt"
)

func main() {
    // 1. 输出Hello World!
    fmt.Println("Hello World!")

    // 2. 计算1+2的结果并输出
    sum := 1 + 2
    fmt.Println(sum)

    // 3. 声明一个字符串变量并输出
    str := "Golang"
    fmt.Println(str)

    // 4. 声明一个整数变量并输出
    num := 10
    fmt.Println(num)

    // 5. 循环打印1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 6. 声明一个切片并输出
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println(slice)

    // 7. 声明一个结构体并输出
    type Person struct {
        Name string
        Age  int
    }

    p := Person{Name: "Alice", Age: 25}
    fmt.Println(p)
}