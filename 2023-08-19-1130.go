package main

import "fmt"

func main() {
    // 练习1：输出 "Hello, World!"
    fmt.Println("Hello, World!")

    // 练习2：输出 1 到 10 的平方数
    for i := 1; i <= 10; i++ {
        fmt.Println(i * i)
    }

    // 练习3：计算两个数的和，并输出结果
    num1 := 10
    num2 := 20
    sum := num1 + num2
    fmt.Println(sum)

    // 练习4：使用 if-else 判断一个数是奇数还是偶数，并输出结果
    number := 7
    if number%2 == 0 {
        fmt.Println("偶数")
    } else {
        fmt.Println("奇数")
    }

    // 练习5：创建一个切片，包含 1 到 5 的数字，并遍历输出
    numbers := []int{1, 2, 3, 4, 5}
    for _, num := range numbers {
        fmt.Println(num)
    }
}