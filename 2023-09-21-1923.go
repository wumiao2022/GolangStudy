package main

import "fmt"

func main() {
    // 练习1: 输出Hello World!
    fmt.Println("Hello World!")

    // 练习2: 变量声明与赋值
    var x int
    x = 10
    var y float64 = 3.14
    fmt.Println(x, y)

    // 练习3: 算术运算
    a := 5
    b := 7
    sum := a + b
    diff := a - b
    fmt.Println(sum, diff)

    // 练习4: 条件语句
    num := 12
    if num > 10 {
        fmt.Println("大于10")
    } else {
        fmt.Println("小于等于10")
    }

    // 练习5: 循环语句
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}