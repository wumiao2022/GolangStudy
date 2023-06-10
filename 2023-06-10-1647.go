package main

import (
    "fmt"
)

func main() {
    // 1. 求1到100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println(sum)

    // 2. 输出1到100中的奇数，偶数
    for i := 1; i <= 100; i++ {
        if i%2 == 0 {
            fmt.Printf("%d是偶数\n", i)
        } else {
            fmt.Printf("%d是奇数\n", i)
        }
    }

    // 3. 按照升序输出100以内的质数
    for i := 2; i <= 100; i++ {
        isPrime := true
        for j := 2; j < i; j++ {
            if i%j == 0 {
                isPrime = false
                break
            }
        }
        if isPrime {
            fmt.Println(i)
        }
    }

    // 4. 使用冒泡排序对数组进行排序
    arr := []int{1, 5, 3, 8, 2}
    for i := 0; i < len(arr)-1; i++ {
        for j := i + 1; j < len(arr); j++ {
            if arr[i] > arr[j] {
                arr[i], arr[j] = arr[j], arr[i]
            }
        }
    }
    fmt.Println(arr)
}