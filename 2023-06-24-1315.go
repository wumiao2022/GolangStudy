package main

import (
    "fmt"
)

func main() {
    // 练习1：打印Hello, world!
    fmt.Println("Hello, world!")
    
    // 练习2：计算1到100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println(sum)
    
    // 练习3：交换两个变量的值
    a, b := 10, 20
    a, b = b, a
    fmt.Println(a, b)
    
    // 练习4：判断一个数是否是质数
    num := 11
    prime := true
    for i := 2; i < num; i++ {
        if num % i == 0 {
            prime = false
            break
        }
    }
    if prime {
        fmt.Println(num, "是质数")
    } else {
        fmt.Println(num, "不是质数")
    }
}