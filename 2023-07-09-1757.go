package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：计算两个整数的和并输出结果
    a := 5
    b := 3
    sum := a + b
    fmt.Println("The sum of", a, "and", b, "is", sum)

    // 练习3：判断一个数是否为偶数，并输出结果
    num := 10
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 练习4：从1打印到10（包括10）
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习5：计算1到10的累加和并输出结果
    total := 0
    for i := 1; i <= 10; i++ {
        total += i
    }
    fmt.Println("The sum of numbers from 1 to 10 is", total)
}