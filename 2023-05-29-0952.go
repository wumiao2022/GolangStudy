package main

import "fmt"

func main() {
    // 练习1：打印九九乘法表
    for i := 1; i < 10; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d*%d=%d ", j, i, i*j)
        }
        fmt.Println()
    }

    fmt.Println()

    // 练习2：判断一个数是否是质数
    num := 19
    isPrime := true
    for i := 2; i < num; i++ {
        if num%i == 0 {
            isPrime = false
            break
        }
    }
    if isPrime {
        fmt.Printf("%d是质数\n", num)
    } else {
        fmt.Printf("%d不是质数\n", num)
    }

    fmt.Println()

    // 练习3：冒泡排序
    arr := []int{3, 5, 9, 1, 8, 4, 7, 2, 6}
    for i := 0; i < len(arr)-1; i++ {
        for j := 0; j < len(arr)-1-i; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    fmt.Println(arr)
}