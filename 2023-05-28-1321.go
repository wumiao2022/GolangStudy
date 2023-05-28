package main

import (
    "fmt"
)

func main() {
    // 练习1: 打印出1~10之间的偶数
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }

    // 练习2: 计算1+2+...+100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println(sum)

    // 练习3: 给定一个切片，计算其中所有元素的和
    numbers := []int{2, 4, 6, 8, 10}
    sum = 0
    for _, num := range numbers {
        sum += num
    }
    fmt.Println(sum)

    // 练习4: 实现一个函数，给定一个数字n，返回小于等于n的所有素数
    fmt.Println(getPrimes(20)) // 输出 [2 3 5 7 11 13 17 19]
}

func getPrimes(n int) []int {
    primes := []int{}
    for i := 2; i <= n; i++ {
        flag := true
        for j := 2; j <= i/2; j++ {
            if i%j == 0 {
                flag = false
                break
            }
        }
        if flag {
            primes = append(primes, i)
        }
    }
    return primes
}