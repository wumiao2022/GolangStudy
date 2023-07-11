package main

import "fmt"

func main() {
    // 练习1：打印Hello World
    fmt.Println("Hello World")

    // 练习2：打印1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习3：计算两个数的和
    num1 := 5
    num2 := 10
    sum := num1 + num2
    fmt.Println("Sum:", sum)

    // 练习4：判断一个数是否为偶数
    checkEven := func(num int) bool {
        if num%2 == 0 {
            return true
        }
        return false
    }
    fmt.Println("Is 4 even?", checkEven(4))
    fmt.Println("Is 7 even?", checkEven(7))

    // 练习5：计算斐波那契数列的前n项
    fibonacci := func(n int) {
        a, b := 0, 1
        for i := 0; i < n; i++ {
            fmt.Println(a)
            a, b = b, a+b
        }
    }
    fmt.Println("Fibonacci sequence:")
    fibonacci(10)
}
