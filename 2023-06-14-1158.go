package main

import "fmt"

func main() {
    // 练习1：计算两个整数的和
    num1 := 10
    num2 := 20
    sum := num1 + num2
    fmt.Println("Sum of", num1, "and", num2, "is", sum)

    // 练习2：判断一个数是否为偶数
    num := 7
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 练习3：打印九九乘法表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d*%d=%d ", i, j, i*j)
        }
        fmt.Println()
    }

    // 练习4：排序一个字符串切片
    words := []string{"apple", "cat", "dog", "banana", "zebra"}
    for i := 0; i < len(words)-1; i++ {
        for j := 0; j < len(words)-i-1; j++ {
            if words[j] > words[j+1] {
                temp := words[j]
                words[j] = words[j+1]
                words[j+1] = temp
            }
        }
    }
    fmt.Println("Sorted words:", words)
}