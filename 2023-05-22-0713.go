package main

import "fmt"

func main() {
    // 练习1：打印出10以内的奇数
    for i := 1; i <= 10; i++ {
        if i%2 != 0 {
            fmt.Println(i)
        }
    }

    // 练习2：求1到100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println(sum)

    // 练习3：判断一个数是否为素数
    num := 11
    isPrime := true
    for i := 2; i <= num/2; i++ {
        if num%i == 0 {
            isPrime = false
            break
        }
    }
    if isPrime {
        fmt.Println(num, "是素数")
    } else {
        fmt.Println(num, "不是素数")
    }
}