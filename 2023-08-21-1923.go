package main

import "fmt"

func main() {
    // 打印Hello, World!
    fmt.Println("Hello, World!")

    // 判断一个数是否为偶数
    num := 8
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 计算一个数的阶乘
    fact := 1
    for i := num; i >= 1; i-- {
        fact *= i
    }
    fmt.Println("Factorial of", num, "is", fact)

    // 反转字符串
    str := "Hello, Golang!"
    reversed := ""
    for _, char := range str {
        reversed = string(char) + reversed
    }
    fmt.Println("Reversed string:", reversed)

    // 查找切片中的最大值
    numbers := []int{2, 4, 6, 1, 8, 3}
    max := numbers[0]
    for _, num := range numbers {
        if num > max {
            max = num
        }
    }
    fmt.Println("Max value in the slice:", max)

    // 使用函数实现斐波那契数列
    fmt.Println("Fibonacci series:")
    for i := 0; i < 10; i++ {
        fmt.Println(fibonacci(i))
    }
}

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}