package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：打印1到10
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习3：判断一个数是否为偶数
    num := 6
    if num%2 == 0 {
        fmt.Println("偶数")
    } else {
        fmt.Println("奇数")
    }

    // 练习4：计算两个数的和
    a := 3
    b := 7
    sum := a + b
    fmt.Println("和 =", sum)

    // 练习5：使用数组存储一组学生的成绩，并计算平均分
    grades := [5]float64{85.5, 92.0, 78.5, 90.0, 88.5}
    total := 0.0
    for _, grade := range grades {
        total += grade
    }
    average := total / float64(len(grades))
    fmt.Println("平均分 =", average)
}