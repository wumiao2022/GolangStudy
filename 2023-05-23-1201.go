package main

import (
    "fmt"
)

func main() {
    // 练习1：输出Hello World
    fmt.Println("Hello World")

    // 练习2：定义变量并输出
    var name string = "Tom"
    age := 20
    fmt.Println(name, "is", age, "years old")

    // 练习3：循环输出1~10
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习4：切片操作
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println(nums[1:3])

    // 练习5：函数调用
    result := add(2, 3)
    fmt.Println("The result is", result)
}

func add(a, b int) int {
    return a + b
}