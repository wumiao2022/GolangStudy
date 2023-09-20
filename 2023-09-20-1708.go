package main

import "fmt"

func main() {
    // 练习1：打印数字1到10
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习2：计算1到100之间的偶数和
    sum := 0
    for i := 1; i <= 100; i++ {
        if i % 2 == 0 {
            sum += i
        }
    }
    fmt.Println("偶数和:", sum)

    // 练习3：判断是否为素数
    num := 17
    isPrime := true
    for i := 2; i < num; i++ {
        if num%i == 0 {
            isPrime = false
            break
        }
    }
    fmt.Println(num, "是素数吗？", isPrime)

    // 练习4：打印乘法口诀表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d * %d = %d\t", j, i, j*i)
        }
        fmt.Println()
    }
}