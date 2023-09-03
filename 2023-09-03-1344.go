package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：打印1到10的整数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习3：求一个整数数组的和
    nums := []int{1, 2, 3, 4, 5}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("Sum:", sum)

    // 练习4：计算两个数的最大公约数
    a, b := 36, 48
    for b != 0 {
        a, b = b, a%b
    }
    fmt.Println("GCD:", a)

    // 练习5：反转字符串
    str := "Hello, World!"
    reversed := ""
    for _, char := range str {
        reversed = string(char) + reversed
    }
    fmt.Println("Reversed:", reversed)
}