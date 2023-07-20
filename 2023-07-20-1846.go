package main

import "fmt"

func main() {
    // 练习1: 打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2: 计算1到10的和并打印结果
    sum := 0
    for i := 1; i <= 10; i++ {
        sum += i
    }
    fmt.Println("Sum:", sum)

    // 练习3: 判断一个数是否为偶数并打印结果
    num := 5
    if num%2 == 0 {
        fmt.Println("Even")
    } else {
        fmt.Println("Odd")
    }

    // 练习4: 打印1到100之间的所有素数
    for i := 1; i <= 100; i++ {
        prime := true
        for j := 2; j <= i/2; j++ {
            if i%j == 0 {
                prime = false
                break
            }
        }
        if prime {
            fmt.Println(i)
        }
    }

    // 练习5: 创建一个Slice，并使用循环打印其中的元素
    nums := []int{1, 2, 3, 4, 5}
    for _, num := range nums {
        fmt.Println(num)
    }
}