package main

import "fmt"

func main() {
    // 练习1：输出Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：计算并输出1+2的结果
    fmt.Println(1 + 2)

    // 练习3：计算并输出5的阶乘
    fmt.Println(factorial(5))

    // 练习4：将字符串“Golang”反转并输出
    fmt.Println(reverseString("Golang"))
}

// 阶乘函数
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

// 字符串反转函数
func reverseString(str string) string {
    runes := []rune(str)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}