package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano()) // 随机数种子

    // 生成10个随机数
    var nums [10]int
    for i := 0; i < 10; i++ {
        nums[i] = rand.Intn(100)
    }

    fmt.Println("随机数列表：", nums)

    // 选择排序
    for i := 0; i < 10-1; i++ {
        minIndex := i
        for j := i + 1; j < 10; j++ {
            if nums[j] < nums[minIndex] {
                minIndex = j
            }
        }
        nums[i], nums[minIndex] = nums[minIndex], nums[i]
    }

    fmt.Println("排序后的列表：", nums)
}