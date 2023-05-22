package main

import (
    "fmt"
)

func main() {
    // 练习1：输出 "Hello, world!"
    fmt.Println("Hello, world!")
    
    // 练习2：输出整数1到10
    for i:=1; i<=10; i++ {
        fmt.Println(i)
    }
    
    // 练习3：计算1到100的和
    sum := 0
    for i:=1; i<=100; i++ {
        sum += i
    }
    fmt.Println(sum)
    
    // 练习4：判断闰年
    year := 2020
    if year%4 == 0 && year%100 != 0 || year%400 == 0 {
        fmt.Printf("%d是闰年\n", year)
    } else {
        fmt.Printf("%d不是闰年\n", year)
    }
    
    // 练习5：计算圆的面积和周长
    r := 5.0
    area := 3.14 * r * r
    circumference := 2 * 3.14 * r
    fmt.Printf("半径为%.2f的圆的面积为%.2f，周长为%.2f\n", r, area, circumference)
}