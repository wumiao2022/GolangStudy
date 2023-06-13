package main

import "fmt"

func main() {
    // 1. 找出切片中的最大值和最小值
    nums := []int{2, 8, 5, 1, 10, 7}
    max := nums[0]
    min := nums[0]
    for _, num := range nums {
        if num > max {
            max = num
        }
        if num < min {
            min = num
        }
    }
    fmt.Printf("最大值是：%d，最小值是：%d\n", max, min)

    // 2. 删除切片中的元素
    fruits := []string{"apple", "banana", "orange", "grape", "peach"}
    index := 2
    fmt.Println(fruits)
    fruits = append(fruits[:index], fruits[index+1:]...)
    fmt.Println(fruits)

    // 3. 循环输出乘法口诀表
    for i := 1; i <= 9; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d*%d=%d\t", j, i, i*j)
        }
        fmt.Println()
    }
}