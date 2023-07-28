package main

import "fmt"

func main() {
    // 练习1：打印1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习2：判断一个数是奇数还是偶数
    number := 13
    if number%2 == 0 {
        fmt.Println("偶数")
    } else {
        fmt.Println("奇数")
    }

    // 练习3：计算1到100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println(sum)

    // 练习4：将字符串中的大写字母转为小写字母
    str := "GoLANG"
    lowerStr := ""
    for _, char := range str {
        if char >= 'A' && char <= 'Z' {
            lowerChar := char + 32
            lowerStr += string(lowerChar)
        } else {
            lowerStr += string(char)
        }
    }
    fmt.Println(lowerStr)
}