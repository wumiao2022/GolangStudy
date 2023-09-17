package main

import "fmt"

func main() {
    // 实例1：打印Hello World
    fmt.Println("Hello World!")

    // 实例2：计算两个数的和
    a := 10
    b := 20
    sum := a + b
    fmt.Println("Sum:", sum)

    // 实例3：判断一个数是否为偶数
    num := 15
    if num%2 == 0 {
        fmt.Println("Even number")
    } else {
        fmt.Println("Odd number")
    }

    // 实例4：判断一个年份是否为闰年
    year := 2020
    if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
        fmt.Println(year, "is a leap year")
    } else {
        fmt.Println(year, "is not a leap year")
    }

    // 实例5：求一个数的阶乘
    n := 5
    factorial := 1
    for i := 1; i <= n; i++ {
        factorial *= i
    }
    fmt.Println("Factorial of", n, "is", factorial)
}