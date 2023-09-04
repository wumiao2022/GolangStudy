package main

import (
    "fmt"
    "time"
)

func main() {
    // 1. 打印当前时间的小时、分钟和秒数
    now := time.Now()
    fmt.Println("当前时间：", now.Hour(), "时", now.Minute(), "分", now.Second(), "秒")

    // 2. 使用for循环计算1到100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("1到100的和为：", sum)

    // 3. 定义一个切片并进行遍历输出
    fruits := []string{"apple", "banana", "orange"}
    for _, fruit := range fruits {
        fmt.Println("水果：", fruit)
    }

    // 4. 计算两个数的最大公约数
    num1, num2 := 24, 36
    for num2 != 0 {
        temp := num1 % num2
        num1 = num2
        num2 = temp
    }
    fmt.Println("24和36的最大公约数为：", num1)

    // 5. 使用switch语句进行成绩等级判断
    score := 85
    switch {
    case score >= 90:
        fmt.Println("成绩等级为：A")
    case score >= 80:
        fmt.Println("成绩等级为：B")
    case score >= 70:
        fmt.Println("成绩等级为：C")
    case score >= 60:
        fmt.Println("成绩等级为：D")
    default:
        fmt.Println("成绩等级为：E")
    }
}
```