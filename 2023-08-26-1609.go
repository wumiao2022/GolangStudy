package main

import (
    "fmt"
)

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：计算两个整数的和
    var num1, num2 int
    num1 = 10
    num2 = 5
    sum := num1 + num2
    fmt.Println("Sum:", sum)

    // 练习3：判断一个数是否是偶数
    number := 7
    if number%2 == 0 {
        fmt.Println(number, "is even")
    } else {
        fmt.Println(number, "is odd")
    }

    // 练习4：使用循环输出1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习5：使用切片操作删除切片中的元素
    numbers := []int{1, 2, 3, 4, 5}
    indexToDelete := 2
    numbers = append(numbers[:indexToDelete], numbers[indexToDelete+1:]...)
    fmt.Println(numbers)
}