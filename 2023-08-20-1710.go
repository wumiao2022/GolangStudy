package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：求和函数
    sum := add(3, 5)
    fmt.Println("Sum:", sum)

    // 练习3：判断是否为偶数
    num := 10
    if isEven(num) {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 练习4：判断闰年
    year := 2020
    if isLeapYear(year) {
        fmt.Println(year, "is a leap year")
    } else {
        fmt.Println(year, "is not a leap year")
    }
}

func add(a, b int) int {
    return a + b
}

func isEven(num int) bool {
    if num%2 == 0 {
        return true
    }
    return false
}

func isLeapYear(year int) bool {
    if year%400 == 0 {
        return true
    }
    if year%100 == 0 {
        return false
    }
    if year%4 == 0 {
        return true
    }
    return false
}