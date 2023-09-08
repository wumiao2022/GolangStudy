package main

import "fmt"

func main() {
    // 练习1：打印出1-100之间的所有偶数
    for i := 1; i <= 100; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }

    // 练习2：求1-100之间所有奇数的和
    sum := 0
    for i := 1; i <= 100; i++ {
        if i%2 != 0 {
            sum += i
        }
    }
    fmt.Println("Sum of odd numbers:", sum)

    // 练习3：计算阶乘
    num := 5
    factorial := 1
    for i := 1; i <= num; i++ {
        factorial *= i
    }
    fmt.Printf("Factorial of %d is %d\n", num, factorial)
}