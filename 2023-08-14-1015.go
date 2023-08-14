package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：声明一个整数变量并打印其值
    var num int = 10
    fmt.Println(num)

    // 练习3：声明一个字符串变量并打印其值
    var str string = "Hello, Golang!"
    fmt.Println(str)

    // 练习4：声明一个浮点数变量并打印其值
    var f float64 = 3.14
    fmt.Println(f)

    // 练习5：使用if-else语句判断一个数是否为偶数，并打印结果
    num = 7
    if num%2 == 0 {
        fmt.Println("偶数")
    } else {
        fmt.Println("奇数")
    }

    // 练习6：声明一个切片，并将切片中的元素打印出来
    slice := []int{1, 2, 3, 4, 5}
    for _, value := range slice {
        fmt.Println(value)
    }
}