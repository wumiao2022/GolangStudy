package main

import "fmt"

func main() {
    // 实现一个从1加到100的函数
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println(sum)

    // 实现一个打印九九乘法表的函数
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d * %d = %d ", j, i, j*i)
        }
        fmt.Println()
    }

    // 实现一个找出切片中最大值的函数
    nums := []int{3, 7, 2, 9, 4, 1, 6, 8, 5}
    max := nums[0]
    for _, num := range nums {
        if num > max {
            max = num
        }
    }
    fmt.Println(max)
}