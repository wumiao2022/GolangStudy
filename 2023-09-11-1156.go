package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：计算两个数的和并打印结果
    var num1, num2 int = 10, 20
    sum := num1 + num2
    fmt.Println("Sum:", sum)

    // 练习3：判断一个数是否为偶数并打印结果
    num := 15
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 练习4：使用for循环打印1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习5：使用range遍历数组并打印每个元素的值
    arr := []int{1, 2, 3, 4, 5}
    for _, val := range arr {
        fmt.Println(val)
    }
}
