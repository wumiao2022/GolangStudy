package main

import (
    "fmt"
)

func main() {
    // 练习1：输出hello world
    fmt.Println("Hello World")

    // 练习2：计算1~100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println(sum)

    // 练习3：计算1~100的偶数之和
    sumEven := 0
    for i := 1; i <= 100; i++ {
        if i%2 == 0 {
            sumEven += i
        }
    }
    fmt.Println(sumEven)

    // 练习4：数组遍历
    arr := [5]int{1, 2, 3, 4, 5}
    for i, v := range arr {
        fmt.Printf("arr[%d] = %d\n", i, v)
    }

    // 练习5：判断一个数是否为质数
    num := 17
    isPrime := true
    for i := 2; i <= num/2; i++ {
        if num%i == 0 {
            isPrime = false
            break
        }
    }
    if isPrime {
        fmt.Printf("%d is a prime number\n", num)
    } else {
        fmt.Printf("%d is not a prime number\n", num)
    }
}