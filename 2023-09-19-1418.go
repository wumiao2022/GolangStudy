package main


import "fmt"


func main() {
    // 练习1：打印Hello World
    fmt.Println("Hello World")

    // 练习2：计算两个数的和
    num1 := 10
    num2 := 5
    sum := num1 + num2
    fmt.Println("Sum:", sum)

    // 练习3：获取用户输入的姓名并打印
    var name string
    fmt.Print("Please enter your name: ")
    fmt.Scanln(&name)
    fmt.Println("Hello", name)

    // 练习4：使用循环打印1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}