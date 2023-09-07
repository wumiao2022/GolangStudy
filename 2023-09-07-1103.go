package main

import "fmt"

func main() {
    // 练习1：将字符串反转
    str := "Hello, world!"
    reversed := ""
    for i := len(str)-1; i >= 0; i-- {
        reversed += string(str[i])
    }
    fmt.Println(reversed)

    // 练习2：计算字符串中的字母个数
    str = "Hello, world!"
    letterCount := 0
    for _, char := range str {
        if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
            letterCount++
        }
    }
    fmt.Println(letterCount)

    // 练习3：判断一个数是否为质数
    number := 13
    isPrime := true
    for i := 2; i <= number/2; i++ {
        if number%i == 0 {
            isPrime = false
            break
        }
    }
    fmt.Println(isPrime)
}