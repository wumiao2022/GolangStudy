package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano()) // 设置随机数种子

    // 生成随机数数组
    var arr [10]int
    for i := 0; i < len(arr); i++ {
        arr[i] = rand.Intn(100)
    }

    fmt.Println("原数组：", arr)

    // 冒泡排序
    for i := 0; i < len(arr)-1; i++ {
        for j := 0; j < len(arr)-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }

    fmt.Println("排序后的数组：", arr)
}