package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")
    
    // 练习2：定义变量并输出
    var name string = "John"
    fmt.Println("My name is", name)
    
    // 练习3：定义切片并遍历输出
    numbers := []int{1, 2, 3, 4, 5}
    for _, num := range numbers {
        fmt.Println(num)
    }
    
    // 练习4：定义函数并调用
    result := add(2, 3)
    fmt.Println("The addition result is", result)
}

func add(a, b int) int {
    return a + b
}