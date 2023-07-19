package main

import (
    "fmt"
)

func main() {
    // 1. 将两个整数相加并输出结果
    var num1, num2 int = 10, 5
    result := num1 + num2
    fmt.Println("Result:", result)

    // 2. 将一个字符串转换为整数并输出结果
    str := "123"
    integerValue := 0
    for _, char := range str {
        integerValue = integerValue*10 + int(char-'0')
    }
    fmt.Println("Integer Value:", integerValue)

    // 3. 判断一个数是否为质数并输出结果
    number := 17
    isPrime := true
    for i := 2; i <= number/2; i++ {
        if number%i == 0 {
            isPrime = false
            break
        }
    }
    fmt.Println("Is Prime Number:", isPrime)

    // 4. 判断一个字符串是否为回文字符串并输出结果
    text := "level"
    isPalindrome := true
    for i := 0; i < len(text)/2; i++ {
        if text[i] != text[len(text)-1-i] {
            isPalindrome = false
            break
        }
    }
    fmt.Println("Is Palindrome:", isPalindrome)
}