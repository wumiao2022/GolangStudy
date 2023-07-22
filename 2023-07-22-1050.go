package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：计算两个整数的和
    a := 3
    b := 5
    sum := a + b
    fmt.Println("The sum of", a, "and", b, "is", sum)

    // 练习3：判断一个数是否为偶数
    num := 7
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 练习4：判断一个年份是否为闰年
    year := 2022
    if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
        fmt.Println(year, "is a leap year")
    } else {
        fmt.Println(year, "is not a leap year")
    }

    // 练习5：计算1到100之间所有偶数的和
    total := 0
    for i := 1; i <= 100; i++ {
        if i%2 == 0 {
            total += i
        }
    }
    fmt.Println("The sum of all even numbers from 1 to 100 is", total)
}