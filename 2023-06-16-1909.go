package main

import "fmt"

func main() {
    // 1. 打印"Hello, world!"
    fmt.Println("Hello, world!")

    // 2. 输出1到10的数字
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    // 3. 输出1到10的偶数
    for i := 2; i <= 10; i += 2 {
        fmt.Println(i)
    }

    // 4. 从命令行接受一个整数，并判断是否为偶数
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }

    // 5. 计算1到1000的和
    sum := 0
    for i := 1; i <= 1000; i++ {
        sum += i
    }
    fmt.Println("Sum of 1 to 1000:", sum)
}