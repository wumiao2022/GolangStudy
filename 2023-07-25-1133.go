package main

import "fmt"

func main() {
    // 练习1：输出"Hello, World!"
    fmt.Println("Hello, World!")

    // 练习2：计算1+1
    fmt.Println(1 + 1)

    // 练习3：定义一个整数变量并输出
    var num int
    num = 10
    fmt.Println(num)

    // 练习4：定义一个字符串变量并输出
    var str string
    str = "Hello, Golang!"
    fmt.Println(str)

    // 练习5：定义一个切片并输出
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println(slice)
}