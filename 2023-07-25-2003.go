package main

import "fmt"

func main() {
    // 练习1：打印Hello World
    fmt.Println("Hello World")

    // 练习2：创建一个整型变量并赋值为10
    num := 10
    fmt.Println(num)

    // 练习3：创建一个字符串变量并赋值为"Golang"
    text := "Golang"
    fmt.Println(text)

    // 练习4：创建一个布尔型变量并赋值为true
    isTrue := true
    fmt.Println(isTrue)

    // 练习5：创建一个整型数组并打印所有元素
    arr := []int{1, 2, 3, 4, 5}
    for _, value := range arr {
        fmt.Println(value)
    }

    // 练习6：创建一个字符串切片并打印所有元素
    slice := []string{"apple", "banana", "cherry"}
    for _, value := range slice {
        fmt.Println(value)
    }
}
