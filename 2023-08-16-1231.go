package main

import "fmt"

func main() {
    // 1. 输出 "Hello, World!"
    fmt.Println("Hello, World!")

    // 2. 定义一个整数变量并赋值为42，然后将其打印出来
    var num int = 42
    fmt.Println(num)

    // 3. 创建一个切片，并用整数1、2、3初始化它，然后将其打印出来
    slice := []int{1, 2, 3}
    fmt.Println(slice)

    // 4. 定义一个函数add，接收两个整数参数并返回它们的和，然后调用该函数并将结果打印出来
    fmt.Println(add(3, 4))

    // 5. 创建一个名为person的结构体，包含姓名和年龄字段，并分别赋值为"Tom"和25，然后将其打印出来
    person := struct {
        name string
        age  int
    }{
        name: "Tom",
        age:  25,
    }
    fmt.Println(person)
}

func add(a, b int) int {
    return a + b
}