package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：计算两个数的和
    a := 10
    b := 20
    sum := a + b
    fmt.Println("Sum:", sum)

    // 练习3：判断一个数是否为偶数
    num := 25
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 练习4：打印1到100之间的所有奇数
    for i := 1; i <= 100; i += 2 {
        fmt.Println(i)
    }

    // 练习5：计算1到100之间的所有偶数的和
    sumEven := 0
    for i := 2; i <= 100; i += 2 {
        sumEven += i
    }
    fmt.Println("Sum of even numbers:", sumEven)
}