package main

import (
    "fmt"
)

func main() {
    // 练习1：输出 1 ~ 100 之间所有的偶数
    for i := 1; i <= 100; i++ {
        if i%2 == 0 {
            fmt.Printf("%d ", i)
        }
    }
    fmt.Println()

    // 练习2：使用冒泡排序法对数组进行升序排序
    arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
    for i := 0; i < len(arr); i++ {
        for j := i + 1; j < len(arr); j++ {
            if arr[j] < arr[i] {
                arr[i], arr[j] = arr[j], arr[i]
            }
        }
    }
    fmt.Println(arr)

    // 练习3：计算斐波那契数列的第 n 项（n 由用户输入）
    var n int
    fmt.Print("请输入一个正整数：")
    fmt.Scanln(&n)
    if n < 1 {
        fmt.Println("输入有误！")
    } else {
        fmt.Printf("斐波那契数列的第 %d 项为：%d\n", n, fibonacci(n))
    }
}

func fibonacci(n int) int {
    if n == 1 || n == 2 {
        return 1
    } else {
        return fibonacci(n-1) + fibonacci(n-2)
    }
}