package main

import "fmt"

func main() {
    // 练习1：打印1到10的所有整数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习2：打印0到100之间能被3整除的所有整数
    for i := 0; i <= 100; i++ {
        if i%3 == 0 {
            fmt.Println(i)
        }
    }

    // 练习3：计算1到100之间所有整数的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("Sum:", sum)

    // 练习4：判断一个数是否是质数
    number := 37
    isPrime := true
    for i := 2; i <= number/2; i++ {
        if number%i == 0 {
            isPrime = false
            break
        }
    }
    if isPrime {
        fmt.Println(number, "is prime.")
    } else {
        fmt.Println(number, "is not prime.")
    }
}
```