package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：计算1+2的结果并输出
    result := 1 + 2
    fmt.Println(result)

    // 练习3：定义一个整数变量num，并将其赋值为10，然后打印出来
    num := 10
    fmt.Println(num)

    // 练习4：定义一个字符串变量name，并将其赋值为"John"，然后打印出来
    name := "John"
    fmt.Println(name)

    // 练习5：定义一个布尔型变量isTrue，并将其赋值为true，然后打印出来
    isTrue := true
    fmt.Println(isTrue)

    // 练习6：定义一个空的切片，并打印其长度和容量
    emptySlice := []int{}
    fmt.Println(len(emptySlice))
    fmt.Println(cap(emptySlice))
}