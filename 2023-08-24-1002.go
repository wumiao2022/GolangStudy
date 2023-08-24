package main

import "fmt"

func main() {
    // 练习1：将两个整数相加并打印结果
    num1 := 10
    num2 := 5
    sum := num1 + num2
    fmt.Println(sum)

    // 练习2：将两个浮点数相乘并打印结果
    float1 := 2.5
    float2 := 3.5
    product := float1 * float2
    fmt.Println(product)

    // 练习3：将两个字符串拼接并打印结果
    str1 := "Hello"
    str2 := "Go!"
    combinedString := str1 + " " + str2
    fmt.Println(combinedString)

    // 练习4：将两个布尔值进行逻辑与操作并打印结果
    bool1 := true
    bool2 := false
    result := bool1 && bool2
    fmt.Println(result)

    // 练习5：将两个整数相除并打印结果
    dividend := 9
    divisor := 2
    quotient := dividend / divisor
    fmt.Println(quotient)
}