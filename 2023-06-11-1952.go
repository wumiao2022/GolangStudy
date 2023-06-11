package main

import "fmt"

func main() {
    // 1. 使用 if 语句判断一个数是否为偶数，如果是偶数输出 true，如果不是偶数输出 false。
    num := 10
    if num%2 == 0 {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }

    // 2. 使用 for 循环输出 1~10 的所有偶数。
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }

    // 3. 定义两个整数变量 a 和 b，使用 switch 语句实现以下功能：
    // 当a+b > 10 输出 "a + b > 10"
    // 当a+b = 10 输出 "a + b = 10"
    // 当a+b < 10 输出 "a + b < 10"
    a := 5
    b := 6
    switch {
    case a+b > 10:
        fmt.Println("a + b > 10")
    case a+b == 10:
        fmt.Println("a + b = 10")
    case a+b < 10:
        fmt.Println("a + b < 10")
    }

    // 4. 定义一个数组 [1,2,3,4,5]，使用 for 循环计算数组中元素的平均值并输出。
    arr := []int{1, 2, 3, 4, 5}
    sum := 0
    for _, num := range arr {
        sum += num
    }
    avg := sum / len(arr)
    fmt.Println(avg)
}