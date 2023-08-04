package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：打印1到10之间的所有偶数
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }

    // 练习3：计算1到100之间的所有偶数的和
    sum := 0
    for i := 1; i <= 100; i++ {
        if i%2 == 0 {
            sum += i
        }
    }
    fmt.Println(sum)

    // 练习4：打印九九乘法表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d * %d = %d\t", j, i, i*j)
        }
        fmt.Println()
    }

    // 练习5：判断一个数字是否为素数
    num := 17
    isPrime := true
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            isPrime = false
            break
        }
    }
    if isPrime {
        fmt.Printf("%d是素数\n", num)
    } else {
        fmt.Printf("%d不是素数\n", num)
    }
}
