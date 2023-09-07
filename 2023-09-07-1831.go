package main

import "fmt"

func main() {
    // 练习1：打印Hello World
    fmt.Println("Hello World")

    // 练习2：计算两个整数的和
    var a = 10
    var b = 20
    sum := a + b
    fmt.Println("Sum:", sum)

    // 练习3：判断一个数是奇数还是偶数
    num := 25
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 练习4：使用for循环打印1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习5：使用switch语句判断今天是星期几
    day := 3
    switch day {
    case 1:
        fmt.Println("Today is Monday")
    case 2:
        fmt.Println("Today is Tuesday")
    case 3:
        fmt.Println("Today is Wednesday")
    case 4:
        fmt.Println("Today is Thursday")
    case 5:
        fmt.Println("Today is Friday")
    case 6:
        fmt.Println("Today is Saturday")
    case 7:
        fmt.Println("Today is Sunday")
    default:
        fmt.Println("Invalid day")
    }
}