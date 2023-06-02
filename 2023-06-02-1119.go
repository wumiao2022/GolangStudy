package main

import "fmt"

func main() {
    // 练习1：打印出10个斐波那契数列的值
    a, b := 0, 1
    for i := 0; i < 10; i++ {
        fmt.Println(a)
        a, b = b, a+b
    }

    // 练习2：打印出1~100之间所有能被3整除的数
    for i := 1; i <= 100; i++ {
        if i%3 == 0 {
            fmt.Println(i)
        }
    }

    // 练习3：打印出翻转后的字符串
    str := "Hello, world!"
    for i := len(str) - 1; i >= 0; i-- {
        fmt.Print(string(str[i]))
    }
}