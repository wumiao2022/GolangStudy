package main

import "fmt"

func main() {
    // 练习1：打印"Hello, World!"
    fmt.Println("Hello, World!")

    // 练习2：声明两个整数变量并求和，将结果打印出来
    var num1, num2 int = 10, 20
    sum := num1 + num2
    fmt.Println("Sum:", sum)

    // 练习3：声明一个字符串变量并打印出其长度
    str := "Hello, Golang!"
    fmt.Println("Length:", len(str))

    // 练习4：声明一个布尔变量并打印出其值
    isTrue := true
    fmt.Println("Is true:", isTrue)

    // 练习5：使用循环打印1-10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习6：使用条件语句判断一个数是否大于10，并打印出结果
    num := 15
    if num > 10 {
        fmt.Println("Number is greater than 10")
    } else {
        fmt.Println("Number is less than or equal to 10")
    }
}