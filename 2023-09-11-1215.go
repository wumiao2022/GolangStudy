package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：求两个整数的和
    num1 := 10
    num2 := 20
    sum := num1 + num2
    fmt.Println("Sum:", sum)

    // 练习3：判断一个数是偶数还是奇数
    num := 23
    if num%2 == 0 {
        fmt.Println("Even")
    } else {
        fmt.Println("Odd")
    }

    // 练习4：判断一个年份是否为闰年
    year := 2024
    if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
        fmt.Println(year, "is a leap year")
    } else {
        fmt.Println(year, "is not a leap year")
    }

    // 练习5：使用for循环计算1到100的和
    sum = 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("Sum of 1 to 100:", sum)
}