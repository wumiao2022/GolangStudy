package main

import "fmt"

func main() {
    // 练习1：打印Hello, World!
    fmt.Println("Hello, World!")

    // 练习2：打印1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习3：计算两个整数的和
    num1 := 10
    num2 := 20
    sum := num1 + num2
    fmt.Println("Sum:", sum)

    // 练习4：判断一个整数是否是偶数
    num := 15
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 练习5：使用函数交换两个变量的值
    x := 5
    y := 10
    swap(&x, &y)
    fmt.Println("x =", x, "y =", y)
}

// swap函数用于交换两个整数的值
func swap(a, b *int) {
    temp := *a
    *a = *b
    *b = temp
}