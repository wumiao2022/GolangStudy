package main

import (
    "fmt"
    "strconv"
)

func main() {
    // 练习1：打印1~100之间的数字
    for i := 1; i <= 100; i++ {
        fmt.Println(i)
    }

    // 练习2：输出1~100之间的奇数
    for i := 1; i <= 100; i++ {
        if i%2 == 1 {
            fmt.Println(i)
        }
    }

    // 练习3：输出1~100之间的偶数
    for i := 1; i <= 100; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }

    // 练习4：输出100~1之间的数字
    for i := 100; i >= 1; i-- {
        fmt.Println(i)
    }

    // 练习5：输出1~100之间的数字，如果是3的倍数输出fizz，
    // 如果是5的倍数输出buzz，如果既是3的倍数又是5的倍数输出fizzbuzz
    for i := 1; i <= 100; i++ {
        if i%3 == 0 && i%5 == 0 {
            fmt.Println("fizzbuzz")
        } else if i%3 == 0 {
            fmt.Println("fizz")
        } else if i%5 == 0 {
            fmt.Println("buzz")
        } else {
            fmt.Println(i)
        }
    }

    // 练习6：将一个字符串转为整数
    str := "123"
    num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(num)
    }

    // 练习7：将一个整数转为字符串
    num = 456
    str = strconv.Itoa(num)
    fmt.Println(str)

    // 练习8：判断一个字符串是否是回文字符串
    str = "abcba"
    for i := 0; i < len(str)/2; i++ {
        if str[i] != str[len(str)-i-1] {
            fmt.Println("不是回文字符串")
            break
        }
        if i == len(str)/2-1 {
            fmt.Println("是回文字符串")
        }
    }
}