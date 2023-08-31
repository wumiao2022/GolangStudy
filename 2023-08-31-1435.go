package main

import "fmt"

func main() {
    // 练习1：打印数字1到10
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习2：计算1到100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("Sum of numbers from 1 to 100:", sum)

    // 练习3：输出斐波那契数列的前10个数字
    fib := []int{0, 1}
    for i := 2; i < 10; i++ {
        fib = append(fib, fib[i-1]+fib[i-2])
    }
    fmt.Println("Fibonacci series:")
    for _, num := range fib {
        fmt.Println(num)
    }

    // 练习4：判断一个数是否为素数
    prime := 17
    isPrime := true
    for i := 2; i <= prime/2; i++ {
        if prime%i == 0 {
            isPrime = false
            break
        }
    }
    if isPrime {
        fmt.Println(prime, "is a prime number")
    } else {
        fmt.Println(prime, "is not a prime number")
    }
}