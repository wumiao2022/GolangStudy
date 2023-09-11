package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：声明变量并赋值，然后打印变量的值
    var num int = 10
    fmt.Println(num)

    // 练习3：使用条件语句判断一个数的正负性，并打印结果
    if num > 0 {
        fmt.Println("正数")
    } else if num < 0 {
        fmt.Println("负数")
    } else {
        fmt.Println("零")
    }

    // 练习4：使用循环打印1到10的整数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习5：使用数组存储一组数据并遍历打印
    arr := [5]int{1, 2, 3, 4, 5}
    for _, value := range arr {
        fmt.Println(value)
    }
}

// 预期输出：
// Hello, World!
// 10
// 正数
// 1
// 2
// 3
// 4
// 5