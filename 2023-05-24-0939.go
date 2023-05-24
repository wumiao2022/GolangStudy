package main

import "fmt"

func main() {
    // Fibonacci数列的前20个数
    a, b := 0, 1
    for i := 0; i < 20; i++ {
        a, b = b, a+b
        fmt.Println(a)
    }

    // 获取一个切片中的最大值和最小值
    nums := []int{2, 4, 7, 1, 10, 5}
    max, min := nums[0], nums[0]
    for _, num := range nums {
        if num > max {
            max = num
        }
        if num < min {
            min = num
        }
    }
    fmt.Println("Max:", max)
    fmt.Println("Min:", min)

    // 将一个字符串反转
    str := "Hello, world!"
    reversed := ""
    for i := len(str) - 1; i >= 0; i-- {
        reversed += string(str[i])
    }
    fmt.Println(reversed)
}