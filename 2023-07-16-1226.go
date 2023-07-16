package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：定义变量并赋值，打印变量值
    var name string = "Alice"
    fmt.Println("My name is", name)

    // 练习3：使用if语句判断变量值，并打印不同的结果
    age := 25
    if age >= 18 {
        fmt.Println("You are an adult.")
    } else {
        fmt.Println("You are underage.")
    }

    // 练习4：创建并使用一个切片
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println("Third element is", numbers[2])

    // 练习5：使用for循环遍历切片并打印每个元素
    for _, num := range numbers {
        fmt.Println(num)
    }

    // 练习6：创建一个函数并调用
    double := func(x int) int {
        return x * 2
    }
    fmt.Println(double(7))

    // 练习7：使用defer关键字延迟执行函数
    defer fmt.Println("This will be printed at the end")
    fmt.Println("This will be printed first")
}
