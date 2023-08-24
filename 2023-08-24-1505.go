package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    // 生成一个随机整数数组
    arr := make([]int, 10)
    for i := 0; i < len(arr); i++ {
        arr[i] = rand.Intn(100)
    }

    // 输出数组内容
    fmt.Println("Original array:", arr)

    // 查找数组中的最大值和最小值
    max, min := arr[0], arr[0]
    for i := 1; i < len(arr); i++ {
        if arr[i] > max {
            max = arr[i]
        }
        if arr[i] < min {
            min = arr[i]
        }
    }
    fmt.Println("Max value:", max)
    fmt.Println("Min value:", min)

    // 对数组进行冒泡排序
    for i := 0; i < len(arr)-1; i++ {
        for j := 0; j < len(arr)-1-i; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }

    // 输出排序后的数组
    fmt.Println("Sorted array:", arr)
}
