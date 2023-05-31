package main

import (
    "fmt"
)

func main() {
    // 练习一：输出Hello, World!
    fmt.Println("Hello, World!")

    // 练习二：输出1到100的数字，如果数字是3的倍数，则输出Fizz；如果数字是5的倍数，则输出Buzz；如果同时是3和5的倍数，则输出FizzBuzz。
    for i := 1; i <= 100; i++ {
        if i%3 == 0 && i%5 == 0 {
            fmt.Println("FizzBuzz")
        } else if i%3 == 0 {
            fmt.Println("Fizz")
        } else if i%5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
    }

    // 练习三：实现一个函数，可以将字符串中的字母大小写互换
    fmt.Println(swapCase("Hello, World!"))

    // 练习四：实现一个函数，判断一个字符串是否为回文字符串
    fmt.Println(isPalindrome("level"))
}

func swapCase(s string) string {
    result := ""
    for _, c := range s {
        if c >= 'a' && c <= 'z' {
            result += string(c - 32)
        } else if c >= 'A' && c <= 'Z' {
            result += string(c + 32)
        } else {
            result += string(c)
        }
    }
    return result
}

func isPalindrome(s string) bool {
    l, r := 0, len(s)-1
    for l < r {
        if s[l] != s[r] {
            return false
        }
        l++
        r--
    }
    return true
}