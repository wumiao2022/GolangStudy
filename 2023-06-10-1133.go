package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // 生成随机数种子
    rand.Seed(time.Now().UnixNano())

    // 生成一个长度为5的随机整数切片
    nums := make([]int, 5)
    for i := 0; i < len(nums); i++ {
        nums[i] = rand.Intn(100)
    }

    // 输出切片数值
    fmt.Println("生成的随机整数切片为：", nums)

    // 对切片进行冒泡排序
    for i := 0; i < len(nums)-1; i++ {
        for j := 0; j < len(nums)-i-1; j++ {
            if nums[j] > nums[j+1] {
                nums[j], nums[j+1] = nums[j+1], nums[j]
            }
        }
    }

    // 输出排序后的切片数值
    fmt.Println("排序后的随机整数切片为：", nums)
}