package main

import (
    "fmt"
)

func main() {
    // 练习1：输出九九乘法表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d*%d=%d ", j, i, i*j)
        }
        fmt.Println()
    }

    // 练习2：求一组整数的平均值
    nums := []int{5, 6, 7, 8, 9}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    avg := float64(sum) / float64(len(nums))
    fmt.Printf("平均值为%.2f\n", avg)

    // 练习3：遍历一个数组，只输出所有偶数
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    for _, num := range arr {
        if num%2 == 0 {
            fmt.Printf("%d是偶数\n", num)
        }
    }
}