package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：计算两个整数的和
    a := 10
    b := 20
    sum := a + b
    fmt.Println("Sum:", sum)

    // 练习3：判断一个数是否为偶数
    num := 7
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 练习4：使用循环打印前10个自然数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习5：使用函数计算给定数组的总和
    nums := []int{1, 2, 3, 4, 5}
    sum := calculateSum(nums)
    fmt.Println("Sum:", sum)
}

func calculateSum(nums []int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return sum
}