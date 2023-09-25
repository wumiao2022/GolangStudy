package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：计算1到100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("1到100的和为:", sum)

    // 练习3：判断一个数是否为素数
    num := 17
    isPrime := true
    for i := 2; i <= num/2; i++ {
        if num%i == 0 {
            isPrime = false
            break
        }
    }
    if isPrime {
        fmt.Println(num, "是素数")
    } else {
        fmt.Println(num, "不是素数")
    }

    // 练习4：统计一个字符串中的字符和数字个数
    str := "Hello1234"
    charCount := 0
    digitCount := 0
    for _, c := range str {
        if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
            charCount++
        } else if c >= '0' && c <= '9' {
            digitCount++
        }
    }
    fmt.Println("字符个数:", charCount)
    fmt.Println("数字个数:", digitCount)
}