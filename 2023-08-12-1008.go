package main

import "fmt"

func main() {
    // 练习1：打印出从1到10的所有整数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习2：计算1到100的所有整数的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("Sum of numbers from 1 to 100:", sum)

    // 练习3：使用数组存储一组字符串，然后打印出每个字符串
    strings := [5]string{"Hello", "World", "Go", "Programming", "Language"}
    for _, str := range strings {
        fmt.Println(str)
    }
}