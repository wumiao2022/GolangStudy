package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano()) // 设置随机数种子

    // 生成随机数切片
    nums := make([]int, 10)
    for i := range nums {
        nums[i] = rand.Intn(100)
    }
    fmt.Println("原始数据：", nums)

    // 冒泡排序
    for i := len(nums) - 1; i > 0; i-- {
        for j := 0; j < i; j++ {
            if nums[j] > nums[j+1] {
                nums[j], nums[j+1] = nums[j+1], nums[j]
            }
        }
    }
    fmt.Println("冒泡排序后：", nums)

    // 选择排序
    for i := 0; i < len(nums)-1; i++ {
        min := i
        for j := i + 1; j < len(nums); j++ {
            if nums[j] < nums[min] {
                min = j
            }
        }
        nums[i], nums[min] = nums[min], nums[i]
    }
    fmt.Println("选择排序后：", nums)

    // 插入排序
    for i := 1; i < len(nums); i++ {
        j := i
        for j > 0 && nums[j] < nums[j-1] {
            nums[j], nums[j-1] = nums[j-1], nums[j]
            j--
        }
    }
    fmt.Println("插入排序后：", nums)
}