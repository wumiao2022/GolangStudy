package main

import "fmt"

func main() {
    // 练习1：输出Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：输出1到10的平方数
    for i := 1; i <= 10; i++ {
        fmt.Println(i * i)
    }

    // 练习3：计算1到100的和
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println(sum)

    // 练习4：判断一个数字是否为偶数
    num := 8
    if num % 2 == 0 {
        fmt.Println("偶数")
    } else {
        fmt.Println("奇数")
    }

    // 练习5：交换两个变量的值
    a := 10
    b := 20
    a, b = b, a
    fmt.Println("a =", a)
    fmt.Println("b =", b)
}
```