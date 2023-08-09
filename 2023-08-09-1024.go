package main

import "fmt"

func main() {
    // 练习一：打印 Hello, World!
    fmt.Println("Hello, World!")

    // 练习二：输出 1-10 的所有整数
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 练习三：计算两个整数的和并输出结果
    a, b := 10, 5
    sum := a + b
    fmt.Println("Sum:", sum)

    // 练习四：声明一个字符串切片，并输出切片长度和第三个元素的值
    names := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
    fmt.Println("Length:", len(names))
    fmt.Println("Third Element:", names[2])

    // 练习五：使用函数交换两个变量的值，然后输出交换后的结果
    x, y := 10, 20
    swap(&x, &y)
    fmt.Println("x:", x, "y:", y)
}

func swap(a, b *int) {
    temp := *a
    *a = *b
    *b = temp
}
