package main

import "fmt"

func main() {
    // 练习1：计算1~100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println(sum)

    // 练习2：打印1~1000中的所有质数
    for i := 2; i <= 1000; i++ {
        isPrime := true
        for j := 2; j < i; j++ {
            if i%j == 0 {
                isPrime = false
                break
            }
        }
        if isPrime {
            fmt.Println(i)
        }
    }

    // 练习3：求出数组[1, 7, 3, 4, 6, 2, 5]的平均值
    arr := [7]int{1, 7, 3, 4, 6, 2, 5}
    sum = 0
    for _, num := range arr {
        sum += num
    }
    avg := float64(sum) / float64(len(arr))
    fmt.Println(avg)
    
    // 练习4：交换两个变量的值
    a, b := 1, 2
    a, b = b, a
    fmt.Println(a, b)

    // 练习5：计算9的阶乘
    factorial := 1
    for i := 1; i <= 9; i++ {
        factorial *= i
    }
    fmt.Println(factorial)
}