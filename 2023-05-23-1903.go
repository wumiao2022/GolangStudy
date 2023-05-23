package main

import "fmt"

func main() {
    // 实例1：打印Hello, world!
    fmt.Println("Hello, world!")
    
    // 实例2：计算并打印1+2的结果
    fmt.Println(1 + 2)
    
    // 实例3：创建一个名为message的字符串变量并初始化
    message := "Hello, Golang!"
    fmt.Println(message)
    
    // 实例4：创建一个名为numbers的整数切片并初始化，长度为3，容量为5
    numbers := make([]int, 3, 5)
    fmt.Println(numbers)
    
    // 实例5：创建一个名为user的结构体类型，包含name和age字段
    type user struct {
        name string
        age int
    }
    var u user
    fmt.Println(u)
}