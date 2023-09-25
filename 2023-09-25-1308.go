package main

import "fmt"

func main() {
    // 练习1：打印九九乘法表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d * %d = %d\t", j, i, j*i)
        }
        fmt.Println()
    }

    fmt.Println("--------------------------")

    // 练习2：求一个整数数组的和
    nums := []int{1, 2, 3, 4, 5}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("Sum of array:", sum)

    fmt.Println("--------------------------")

    // 练习3：反转字符串
    str := "Hello, World!"
    runes := []rune(str)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    fmt.Println("Reversed string:", string(runes))

    fmt.Println("--------------------------")

    // 练习4：判断一个字符串是否是回文字符串
    word := "level"
    isPalindrome := true
    for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
        if word[i] != word[j] {
            isPalindrome = false
            break
        }
    }
    fmt.Println("Is palindrome:", isPalindrome)
}