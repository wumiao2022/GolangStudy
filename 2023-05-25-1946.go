package main

import "fmt"

func main() {
    // 1. 计算两个数的和
    num1 := 10
    num2 := 20
    sum := num1 + num2
    fmt.Println(sum)

    // 2. 计算两个数的差
    num3 := 30
    num4 := 15
    diff := num3 - num4
    fmt.Println(diff)

    // 3. 计算两个数的积
    num5 := 5
    num6 := 6
    product := num5 * num6
    fmt.Println(product)

    // 4. 计算两个数的商
    num7 := 20.0
    num8 := 6.0
    quotient := num7 / num8
    fmt.Println(quotient)

    // 5. 计算一个数的平方
    num9 := 5
    square := num9 * num9
    fmt.Println(square)

    // 6. 计算一个数的立方
    num10 := 3
    cube := num10 * num10 * num10
    fmt.Println(cube)

    // 7. 转换华氏温度为摄氏温度
    fahrenheit := 68.0
    celsius := (fahrenheit - 32) * 5 / 9
    fmt.Println(celsius)

    // 8. 计算一个三角形的面积
    base := 3.0
    height := 4.0
    area := 0.5 * base * height
    fmt.Println(area)

    // 9. 判断一个数是否为偶数
    num11 := 4
    if num11 % 2 == 0 {
        fmt.Println("偶数")
    } else {
        fmt.Println("奇数")
    }

    // 10. 判断一个数是否为质数
    num12 := 13
    isPrime := true
    for i := 2; i < num12 - 1; i++ {
        if num12 % i == 0 {
            isPrime = false
            break
        }
    }
    if isPrime {
        fmt.Println("质数")
    } else {
        fmt.Println("非质数")
    }
}