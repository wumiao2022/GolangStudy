package main

import "fmt"

func main() {
    // 练习一：输出1~10之间的奇数
    for i := 1; i <= 10; i++ {
        if i%2 != 0 {
            fmt.Println(i)
        }
    }

    // 练习二：实现一个简单的加法计算器
    var num1, num2 int
    fmt.Println("请输入两个整数：")
    fmt.Scanln(&num1, &num2)
    fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)

    // 练习三：使用递归函数计算斐波那契数列第n项的值
    fmt.Println(fibonacci(10))
}

func fibonacci(n int) int {
    if n == 0 || n == 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}