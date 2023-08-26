package main

import "fmt"

func main() {
    // 练习1：计算两个整数的和
    var num1, num2, sum int
    num1 = 10
    num2 = 20
    sum = num1 + num2
    fmt.Println("Sum:", sum)

    // 练习2：判断一个数是否为奇数
    number := 15
    if number%2 == 0 {
        fmt.Println(number, "is even")
    } else {
        fmt.Println(number, "is odd")
    }

    // 练习3：打印1到10之间的所有偶数
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            fmt.Println(i)
        }
    }

    // 练习4：使用递归计算阶乘
    factorial := func(n int) int {
        if n == 0 {
            return 1
        }
        return n * factorial(n-1)
    }
    fmt.Println("Factorial of 5:", factorial(5))
}