package main

import "fmt"

func main() {
    // 1. 打印Hello, World!
    fmt.Println("Hello, World!")

    // 2. 输出1到10之间的所有整数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
    
    // 3. 计算1到100之间所有整数的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("Sum:", sum)

    // 4. 判断一个数是否为偶数
    num := 5
    if num % 2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 5. 求一个列表的平均值
    list := []int{1, 2, 3, 4, 5}
    total := 0
    for _, num := range list {
        total += num
    }
    average := float64(total) / float64(len(list))
    fmt.Println("Average:", average)
}

// 以上代码片段包含五个练习实例，分别是打印Hello, World!、输出1到10之间的所有整数、计算1到100之间所有整数的和、判断一个数是否为偶数、求一个列表的平均值。