package main

import "fmt"

func main() {
    // 1. 打印 "Hello, World!"
    fmt.Println("Hello, World!")

    // 2. 声明并初始化一个整数变量x，并打印其值
    x := 10
    fmt.Println(x)

    // 3. 声明并初始化一个浮点数变量y，并打印其值
    y := 3.14
    fmt.Println(y)

    // 4. 声明并初始化一个字符串变量name，并打印其值
    name := "Gopher"
    fmt.Println(name)

    // 5. 声明一个变量a为5，b为3，计算a与b的和并打印结果
    a := 5
    b := 3
    sum := a + b
    fmt.Println(sum)

    // 6. 声明一个变量c为8，计算c的平方并打印结果
    c := 8
    square := c * c
    fmt.Println(square)

    // 7. 声明一个变量d为12，将其转换为浮点数并打印结果
    d := 12
    floatD := float64(d)
    fmt.Println(floatD)

    // 8. 声明一个切片slice1，包含整数1、2、3，并打印其值和长度
    slice1 := []int{1, 2, 3}
    fmt.Println(slice1)
    fmt.Println(len(slice1))
}