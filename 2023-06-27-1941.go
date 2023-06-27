package main

import "fmt"

func main() {
    // 1. 输出Hello, World!
    fmt.Println("Hello, World!")

    // 2. 输出1~100中的偶数
    for i := 1; i <= 100; i++ {
        if i % 2 == 0 {
            fmt.Println(i)
        }
    }

    // 3. 计算1~100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("sum:", sum)

    // 4. 逆序输出"Hello, World!"
    str := "Hello, World!"
    for i := len(str) - 1; i >= 0; i-- {
        fmt.Print(string(str[i]))
    }
    fmt.Println()

    // 5. 判断一个数是否为质数
    num := 17
    isPrime := true
    for i := 2; i < num; i++ {
        if num % i == 0 {
            isPrime = false
            break
        }
    }
    if isPrime {
        fmt.Println(num, "is a prime number")
    } else {
        fmt.Println(num, "is not a prime number")
    }
}