package main

import "fmt"

func main() {
    // 1. 打印Hello, World!
    fmt.Println("Hello, World!")

    // 2. 声明两个变量x和y，分别赋值为10和5，求它们的和、差、乘积和商
    x := 10
    y := 5
    sum := x + y
    diff := x - y
    product := x * y
    quotient := x / y
    fmt.Printf("Sum: %d\nDifference: %d\nProduct: %d\nQuotient: %d\n", sum, diff, product, quotient)

    // 3. 计算1到100的累加和
    total := 0
    for i := 1; i <= 100; i++ {
        total += i
    }
    fmt.Println("Sum of numbers from 1 to 100:", total)

    // 4. 使用条件判断输出一个数是否为偶数或奇数
    num := 17
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 5. 声明一个字符串变量并输出其长度和第一个字符
    str := "Golang"
    strLength := len(str)
    firstChar := str[0]
    fmt.Printf("String length: %d\nFirst character: %c\n", strLength, firstChar)
}