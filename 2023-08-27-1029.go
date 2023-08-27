package main

import (
    "fmt"
)

func main() {
    // 练习1：输出Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：计算两个数的和并输出结果
    a := 5
    b := 3
    sum := a + b
    fmt.Println("The sum of", a, "and", b, "is", sum)

    // 练习3：判断一个数是否为偶数并输出结果
    num := 7
    isEven := num%2 == 0
    fmt.Println(num, "is even:", isEven)

    // 练习4：使用循环输出1到10之间的所有奇数
    for i := 1; i <= 10; i += 2 {
        fmt.Println(i)
    }
}

// 保持学习，每天多练习！