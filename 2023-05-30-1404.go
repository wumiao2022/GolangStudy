package main

import "fmt"

func main() {
    // 练习1
    var num int
    fmt.Println("请输入一个整数：")
    fmt.Scanf("%d", &num)
    if num%2 == 0 {
        fmt.Printf("%d是偶数\n", num)
    } else {
        fmt.Printf("%d是奇数\n", num)
    }

    // 练习2
    var arr = []int{1, 2, 3, 4, 5}
    for i := 0; i < len(arr); i++ {
        fmt.Printf("%d ", arr[i])
    }
    fmt.Println()

    // 练习3
    var str string = "hello"
    var reverse string
    for i := len(str) - 1; i >= 0; i-- {
        reverse += string(str[i])
    }
    fmt.Printf("%s反转后为%s\n", str, reverse)
    
    // 练习4
    var i, j int
    for i = 1; i <= 9; i++ {
        for j = 1; j <= i; j++ {
            fmt.Printf("%dx%d=%d ", j, i, i*j)
        }
        fmt.Println()
    }
}