package main

import "fmt"

func main() {
    // 练习1: 打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2: 计算并打印1到10的和
    sum := 0
    for i := 1; i <= 10; i++ {
        sum += i
    }
    fmt.Println("Sum:", sum)

    // 练习3: 判断一个数是否为偶数
    num := 6
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 练习4: 交换两个变量的值
    a := 5
    b := 10
    a, b = b, a
    fmt.Println("a:", a, "b:", b)

    // 练习5: 根据用户输入打印相应的信息
    var age int
    fmt.Print("请输入您的年龄：")
    fmt.Scan(&age)

    if age >= 18 {
        fmt.Println("您已经成年")
    } else {
        fmt.Println("您尚未成年")
    }
}
