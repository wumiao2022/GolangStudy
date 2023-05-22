package main

import "fmt"

func main() {
    // 1. 打印自己的名字和年龄
    fmt.Println("My name is John, and my age is 30.")

    // 2. 计算并输出两个整数相加的结果
    x := 3
    y := 5
    sum := x + y
    fmt.Println(sum)

    // 3. 输出100以内的所有偶数
    for i := 0; i <= 100; i += 2 {
        fmt.Println(i)
    }

    // 4. 定义一个字符串变量，判断字符串长度是否大于等于5，如果是则输出该字符串，否则输出"字符串长度不够"
    s := "Hello"
    if len(s) >= 5 {
        fmt.Println(s)
    } else {
        fmt.Println("字符串长度不够")
    }

    // 5. 定义一个结构体表示一个人的信息，包括姓名、年龄、性别等字段，并输出其中一个人的信息
    type Person struct {
        Name   string
        Age    int
        Gender string
    }
    person1 := Person{"John", 30, "male"}
    fmt.Println(person1)
}