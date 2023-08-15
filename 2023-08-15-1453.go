package main

import (
    "fmt"
    "time"
)

func main() {
    // 1. 打印当前时间
    fmt.Println(time.Now())

    // 2. 将字符串转换为整数
    numStr := "123"
    num, _ := strconv.Atoi(numStr)
    fmt.Println(num)

    // 3. 将整数转换为字符串
    num := 456
    numStr := strconv.Itoa(num)
    fmt.Println(numStr)

    // 4. 判断一个数是否是偶数
    num := 7
    if num%2 == 0 {
        fmt.Println("Even")
    } else {
        fmt.Println("Odd")
    }

    // 5. 计算两个数的最大公约数
    num1 := 10
    num2 := 15
    for num2 != 0 {
        temp := num2
        num2 = num1 % num2
        num1 = temp
    }
    fmt.Println("GCD:", num1)
}
