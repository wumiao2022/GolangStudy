package main

import "fmt"

func main() {
    // 示例1：打印Hello World
    fmt.Println("Hello World")

    // 示例2：计算两个整数的和
    num1 := 10
    num2 := 20
    sum := num1 + num2
    fmt.Println("Sum:", sum)

    // 示例3：判断一个数是否为偶数
    number := 15
    if number%2 == 0 {
        fmt.Println(number, "is even")
    } else {
        fmt.Println(number, "is odd")
    }

    // 示例4：使用循环打印1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 示例5：定义一个结构体并输出其中的字段
    type Person struct {
        Name    string
        Age     int
        Country string
    }

    person := Person{
        Name:    "John",
        Age:     30,
        Country: "USA",
    }
    fmt.Println("Name:", person.Name)
    fmt.Println("Age:", person.Age)
    fmt.Println("Country:", person.Country)
}
