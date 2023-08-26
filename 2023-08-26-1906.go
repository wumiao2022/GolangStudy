package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：计算两个整数的和并打印结果
    var num1, num2 int
    num1 = 10
    num2 = 5
    sum := num1 + num2
    fmt.Println("Sum:", sum)

    // 练习3：判断一个数是否为偶数并打印结果
    number := 7
    if number%2 == 0 {
        fmt.Println(number, "is even")
    } else {
        fmt.Println(number, "is odd")
    }

    // 练习4：判断一个年份是否为闰年并打印结果
    year := 2024
    if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
        fmt.Println(year, "is a leap year")
    } else {
        fmt.Println(year, "is not a leap year")
    }

    // 练习5：使用循环打印数字1到10
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}

// 练习6：编写一个函数，计算两个整数的乘积并返回结果
func multiply(num1, num2 int) int {
    return num1 * num2
}