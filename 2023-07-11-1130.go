package main

import "fmt"

func main() {
    // 练习1: 输出Hello World
    fmt.Println("Hello World")
    
    // 练习2: 打印数字1到10
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
    
    // 练习3: 求两个数的和
    num1 := 10
    num2 := 5
    sum := num1 + num2
    fmt.Println("Sum:", sum)
    
    // 练习4: 判断一个数是奇数还是偶数
    num := 7
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }
    
    // 练习5: 判断一个年份是否是闰年
    year := 2020
    if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
        fmt.Println(year, "is a leap year")
    } else {
        fmt.Println(year, "is not a leap year")
    }
}