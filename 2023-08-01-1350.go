package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：计算并打印1到100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("Sum:", sum)

    // 练习3：判断一个数是否为素数，并打印结果
    num := 17
    isPrime := true
    for i := 2; i < num; i++ {
        if num%i == 0 {
            isPrime = false
            break
        }
    }
    fmt.Printf("%d is prime: %t\n", num, isPrime)

    // 练习4：编写一个函数，返回两个整数的最大值，并打印结果
    max := getMax(10, 20)
    fmt.Println("Max:", max)
}

func getMax(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
