package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：计算并打印1+2的结果
    result := 1 + 2
    fmt.Println(result)

    // 练习3：使用循环输出1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习4：使用条件语句判断一个数是否为偶数，并打印结果
    num := 8
    if num%2 == 0 {
        fmt.Println("偶数")
    } else {
        fmt.Println("奇数")
    }

    // 练习5：定义一个切片，添加元素后打印整个切片
    var slice []int
    slice = append(slice, 1, 2, 3)
    fmt.Println(slice)
}

// 练习6：实现一个函数，传入两个整数，返回它们的和
func add(x, y int) int {
    return x + y
}