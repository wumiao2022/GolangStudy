package main

import "fmt"

func main() {
    // 练习1: 输出Hello, World!
    fmt.Println("Hello, World!")

    // 练习2: 计算两个整数的和
    var a, b int
    fmt.Print("请输入第一个整数: ")
    fmt.Scan(&a)
    fmt.Print("请输入第二个整数: ")
    fmt.Scan(&b)
    fmt.Printf("两个整数的和为: %d\n", a+b)

    // 练习3: 判断一个数是奇数还是偶数
    var num int
    fmt.Print("请输入一个整数: ")
    fmt.Scan(&num)
    if num%2 == 0 {
        fmt.Println("这个数是偶数")
    } else {
        fmt.Println("这个数是奇数")
    }

    // 练习4: 判断一个年份是否是闰年
    var year int
    fmt.Print("请输入一个年份: ")
    fmt.Scan(&year)
    if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
        fmt.Println("这是一个闰年")
    } else {
        fmt.Println("这不是一个闰年")
    }

    // 练习5: 输出九九乘法表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d * %d = %2d  ", j, i, j*i)
        }
        fmt.Println()
    }
}