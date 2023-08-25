package main

import "fmt"

func main() {
    // 练习1：打印Hello World
    fmt.Println("Hello World")

    // 练习2：打印1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习3：求两个整数的和
    num1 := 10
    num2 := 5
    sum := num1 + num2
    fmt.Println(sum)

    // 练习4：判断一个数是奇数还是偶数
    number := 15
    if number%2 == 0 {
        fmt.Println("偶数")
    } else {
        fmt.Println("奇数")
    }

    // 练习5：使用函数计算斐波那契数列的第n项（n>=0）
    n := 6
    fibonacci := fib(n)
    fmt.Printf("第%d项斐波那契数列为：%d\n", n, fibonacci)
}

func fib(n int) int {
    if n <= 1 {
        return n
    }
    return fib(n-1) + fib(n-2)
}