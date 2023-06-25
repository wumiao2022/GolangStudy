package main

import (
    "fmt"
)

func main() {
    // 练习1：打印出1-100之间所有的偶数
    for i := 1; i <= 100; i++ {
        if i % 2 == 0 {
            fmt.Println(i)
        }
    }
    
    // 练习2：实现一个阶乘函数
    fmt.Println(factorial(5)) // 120
    
    // 练习3：判断一个字符串是否是回文字符串
    fmt.Println(isPalindrome("racecar")) // true
    fmt.Println(isPalindrome("hello")) // false
}

// 阶乘函数
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

// 判断一个字符串是否是回文字符串
func isPalindrome(s string) bool {
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-1-i] {
            return false
        }
    }
    return true
}